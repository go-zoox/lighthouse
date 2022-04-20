package lighthouse

type Config struct {
	Title   string
	Version string
	Author  string
	Server  struct {
		Host string
		Port string
	}
	Cache struct {
		Engine string
		MaxAge string
		Config struct {
			Host     string
			Port     string
			Db       int64
			Password string
		}
	}
	Log struct {
		Transport string
		Level     string
	}
}
