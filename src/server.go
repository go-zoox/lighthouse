package lighthouse

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-zoox/dns"
	"github.com/go-zoox/kv"
	kvtyping "github.com/go-zoox/kv/typing"
	"github.com/go-zoox/logger"
)

// Config is the configuration of lighthouse
func Serve(cfg *Config) {
	server := dns.NewServer(&dns.ServerOptions{
		Port: 53,
	})

	var servers []*dns.ClientDNSServer
	for _, upstream := range cfg.Upstreams {
		servers = append(servers, dns.NewClientDNSServer(upstream, 53))
	}

	client := dns.NewClient(&dns.ClientOptions{
		Servers: servers,
	})

	// @TODO
	if cfg.Cache.Engine == "" {
		cfg.Cache.Engine = "memory"
	}

	cache, err := kv.New(&kvtyping.Config{
		Engine: cfg.Cache.Engine,
		Config: &cfg.Cache.Config,
	})
	if err != nil {
		logger.Error("failed to create cache", err)
		os.Exit(1)
	}

	server.Handle(func(host string, typ int) ([]string, error) {
		if host == "zero.com" && typ == 4 {
			return []string{"6.6.6.6"}, nil
		}

		key := fmt.Sprintf("%s:%d", host, typ)
		if cache.Has(key) {
			ipstr := cache.Get(key)
			var ips []string
			json.Unmarshal([]byte(ipstr), &ips)
			return ips, nil
		}

		if ips, err := client.LookUp(host, &dns.LookUpOptions{Typ: typ}); err != nil {
			ipstr, _ := json.Marshal([]string{})
			cache.Set(key, string(ipstr), 1*60*1000)
			return nil, err
		} else {
			ipstr, _ := json.Marshal(ips)
			cache.Set(key, string(ipstr), 5*60*1000)
			logger.Info("found host(%s %d) %v", host, typ, ips)
			return ips, nil
		}
	})

	server.Serve()
}
