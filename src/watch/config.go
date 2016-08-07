package watch

import "fmt"

const (
	DefaultGrpcPort   = 18858
	DefaultHealthPort = 8080
)

type Config struct {
	Debug         bool
	GrpcPort      int
	HealthPort    int
	TlsCertFile   string
	TlsKeyFile    string
	NoAuth        bool
	ClientID      string
	ClientSecret  string
	AuthDiscovery string

	RedisAddr       string
	RedisPassword   string
	RedisDB         int64
	RedisSentinel   bool
	RedisMasterName string
	NoRedis         bool
}

func (c *Config) GetGrpcPortString() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func (c *Config) GetHealthPortString() string {
	return fmt.Sprintf(":%d", c.HealthPort)
}

func NewConfig() *Config {
	return &Config{GrpcPort: DefaultGrpcPort, RedisMasterName: "mymaster", HealthPort: DefaultHealthPort}
}
