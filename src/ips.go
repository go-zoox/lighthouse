package lighthouse

import (
	"errors"
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

// IPS is the IPs manager
type IPS struct {
	cache  kvtyping.KV
	hosts  *hostsParser.Hosts
	client *dnsClient.Client
}

// NewIPS creates a new IPS
func NewIPS(cfg *Config) *IPS {
	var servers []string = cfg.Upstreams

	client := dns.NewClient(&dns.ClientOptions{
		Servers: servers,
	})

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

	return &IPS{
		cache:  cache,
		hosts:  hosts,
		client: client,
	}
}

// GetByHost returns the IPs of the extract host
func (i *IPS) GetByHost(host string, typ int) ([]string, error) {
	key := i.getKey(host, typ)

	// from cache
	if ips, err := i.getFromCache(key); err == nil {
		return ips, nil
	}

	// from host
	if ips, err := i.getFromHosts(key, host, typ); err == nil {
		return ips, nil
	}

	// from upstream
	if ips, err := i.client.LookUp(host, &dnsClient.LookUpOptions{Typ: typ}); err != nil {
		i.cache.Set(key, []string{}, 1*60*1000)
		return nil, err
	} else {
		i.cache.Set(key, ips, 5*60*1000)
		logger.Info("found host(%s %d) %v", host, typ, ips)
		return ips, nil
	}
}

// GetByWildcardHost returns the IPs of the wildcard host
func (i *IPS) GetByWildcardHost(host string, typ int) ([]string, error) {
	wildcardHosts, _ := GetWildcardHost(host)
	if len(wildcardHosts) > 0 {
		for _, wildcardHost := range wildcardHosts {
			ips, err := i.getByWildHostOne(wildcardHost, typ)
			if err == nil && len(ips) > 0 {
				return ips, nil
			}
		}
	}

	return nil, errors.New("not found wildcard host:" + host)
}

func (i *IPS) getByWildHostOne(host string, typ int) ([]string, error) {
	key := i.getKey(host, typ)

	// from cache
	if ips, err := i.getFromCache(key); err == nil {
		return ips, nil
	}

	// from host
	if ips, err := i.getFromHosts(key, host, typ); err == nil {
		return ips, nil
	}

	return nil, errors.New("not found wildcard host:" + host)
}

func (i *IPS) getKey(host string, typ int) string {
	return fmt.Sprintf("%s:%d", host, typ)
}

func (i *IPS) getFromCache(key string) ([]string, error) {
	// from cache
	if i.cache.Has(key) {
		var ips []string
		err := i.cache.Get(key, &ips)
		if err != nil {
			return nil, err
		}

		return ips, nil
	}

	return nil, errors.New("not found cache key:" + key)
}

func (i *IPS) getFromHosts(key string, host string, typ int) ([]string, error) {
	if i.hosts != nil {
		if ip, err := i.hosts.LookUp(host, typ); err == nil {
			ips := []string{ip}
			i.cache.Set(key, ips, 5*60*1000)
			return ips, nil
		}
	}

	return nil, errors.New("not found hosts key:" + key)
}