package lighthouse

import (
	"github.com/go-zoox/dns"
)

// Config is the configuration of lighthouse
func Serve(cfg *Config) {
	server := dns.NewServer(&dns.ServerOptions{
		Port: 53,
	})

	ipsSearch := NewIPS(cfg)

	server.Handle(func(host string, typ int) ([]string, error) {
		if host == "zero.com" && typ == 4 {
			return []string{"6.6.6.6"}, nil
		}

		if ips, err := ipsSearch.GetByWildcardHost(host, typ); err == nil {
			return ips, nil
		}

		return ipsSearch.GetByHost(host, typ)
	})

	server.Serve()
}
