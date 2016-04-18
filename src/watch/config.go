package watch

import "fmt"

const (
	DefaultGrpcPort = 18858
)

type Config struct {
	Debug         bool
	GrpcPort      int
	TlsCertFile   string
	TlsKeyFile    string
	NoAuth        bool
	ClientID      string
	ClientSecret  string
	AuthDiscovery string

	RedisAddr     string
	RedisPassword string
	RedisDB       int64
	RedisSentinel bool
	NoRedis       bool
}

func (c *Config) GetGrpcPortString() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func NewConfig() *Config {
	return &Config{GrpcPort: DefaultGrpcPort}
}
