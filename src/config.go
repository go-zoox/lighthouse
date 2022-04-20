package lighthouse

import "github.com/go-zoox/kv/redis"

// Config is the configuration of lighthouse
type Config struct {
	Title   string
	Version string
	Author  string
	Server  struct {
		Host string
		Port int64
	}
	Cache struct {
		Engine string
		MaxAge string
		// Config struct {
		// 	Host     string
		// 	Port     int
		// 	Db       int
		// 	Password string
		// 	Prefix   string
		// }
		Config redis.RedisConfig
	}
	Log struct {
		Transport string
		Level     string
	}
}
