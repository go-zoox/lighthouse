package lighthouse

// Config is the configuration of lighthouse
type Config struct {
	Title   string `config:"title"`
	Version string `config:"version"`
	Author  string `config:"author"`
	Server  struct {
		Host string `config:"host"`
		Port int64  `config:"port"`
	} `config:"server"`
	Upstreams []string `config:"upstreams"`
	Cache     struct {
		Engine string `config:"engine,default=memory"`
		MaxAge int64  `config:"max_age"`
		Config struct {
			Host     string `config:"host"`
			Port     int64  `config:"port"`
			Db       int64  `config:"db"`
			Password string `config:"password"`
			Prefix   string `config:"prefix"`
		} `config:"config"`
		// Config redis.RedisConfig `config:"config"`
	} `config:"cache"`
	Log struct {
		Transport string `config:"transport"`
		Level     string `config:"level"`
	} `config:"log"`
	Hosts struct {
		Enable bool   `config:"enable"`
		File   string `config:"file"`
	} `config:"hosts"`
}
