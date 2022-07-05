package core

import (
	"github.com/go-zoox/dns"
)

// Config is the configuration of lighthouse
func Serve(cfg *Config) {
	server := dns.NewServer(&dns.ServerOptions{
		Port: int(cfg.Server.Port),
	})

	ipsSearch := NewIPS(cfg)

	server.Handle(func(host string, typ int) ([]string, error) {
		if host == "zero.com" && typ == 4 {
			return []string{"6.6.6.6"}, nil
		}

		if ips, err := ipsSearch.GetByExactHost(host, typ); err == nil {
			return ips, nil
		}

		if ips, err := ipsSearch.GetByWildcardHost(host, typ); err == nil {
			return ips, nil
		}

		return ipsSearch.SearchByHost(host, typ)
	})

	server.Serve()
}
