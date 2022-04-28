package lighthouse

import (
	"fmt"
	"os"

	"github.com/go-zoox/dns"
	dnsClient "github.com/go-zoox/dns/client"
	"github.com/go-zoox/fs"
	hostsParser "github.com/go-zoox/fs/type/hosts"
	"github.com/go-zoox/kv"
	"github.com/go-zoox/kv/redis"
	kvtyping "github.com/go-zoox/kv/typing"
	"github.com/go-zoox/logger"
)

// Config is the configuration of lighthouse
func Serve(cfg *Config) {
	server := dns.NewServer(&dns.ServerOptions{
		Port: 53,
	})

	var servers []string
	for _, upstream := range cfg.Upstreams {
		// @TODO valid upstream
		servers = append(servers, upstream)
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
		Config: &redis.RedisConfig{
			Host:     cfg.Cache.Config.Host,
			Port:     int(cfg.Cache.Config.Port),
			DB:       int(cfg.Cache.Config.Db),
			Password: cfg.Cache.Config.Password,
			Prefix:   cfg.Cache.Config.Prefix,
		},
	})
	if err != nil {
		logger.Error("failed to create cache", err)
		os.Exit(1)
	}

	var hosts *hostsParser.Hosts
	if cfg.Hosts.Enable {
		if !fs.IsExist(cfg.Hosts.File) {
			logger.Error("hosts file(%s) not found", cfg.Hosts.File)
			os.Exit(1)
		}

		hosts = hostsParser.New(cfg.Hosts.File)
		if err := hosts.Load(); err != nil {
			logger.Error("failed to load hosts file", err)
			os.Exit(1)
		}
	}

	server.Handle(func(host string, typ int) ([]string, error) {
		if host == "zero.com" && typ == 4 {
			return []string{"6.6.6.6"}, nil
		}

		key := fmt.Sprintf("%s:%d", host, typ)

		// from cache
		if cache.Has(key) {
			var ips []string
			err := cache.Get(key, &ips)
			if err != nil {
				return nil, err
			}

			return ips, nil
		}

		// from host
		if hosts != nil {
			if ip, err := hosts.LookUp(host, typ); err == nil {
				ips := []string{ip}
				cache.Set(key, ips, 5*60*1000)
				return ips, nil
			}
		}

		// from upstream
		if ips, err := client.LookUp(host, &dnsClient.LookUpOptions{Typ: typ}); err != nil {
			cache.Set(key, []string{}, 1*60*1000)
			return nil, err
		} else {
			cache.Set(key, ips, 5*60*1000)
			logger.Info("found host(%s %d) %v", host, typ, ips)
			return ips, nil
		}
	})

	server.Serve()
}
