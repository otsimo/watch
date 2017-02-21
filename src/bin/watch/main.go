package main

import (
	"fmt"
	"os"
	"watch"

	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Version string
var config *watch.Config = watch.NewConfig()

func RunAction(c *cli.Context) error {
	config.Debug = c.Bool("debug")
	config.GrpcPort = c.Int("grpc-port")
	config.HealthPort = c.Int("health-port")
	config.TlsCertFile = c.String("tls-cert-file")
	config.TlsKeyFile = c.String("tls-key-file")
	config.ClientID = c.String("client-id")
	config.ClientSecret = c.String("client-secret")
	config.AuthDiscovery = c.String("discovery")
	config.NoAuth = c.Bool("no-auth")
	config.RedisAddr = c.String("redis-addr")
	config.RedisDB = int64(c.Int("redis-db"))
	config.RedisPassword = c.String("redis-password")
	config.RedisMasterName = c.String("redis-master")
	config.RedisSentinel = c.Bool("redis-sentinel")
	config.NoRedis = c.Bool("no-redis")

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	server := watch.NewServer(config)
	return server.ListenGRPC()
}

func withEnvs(prefix string, flags []cli.Flag) []cli.Flag {
	var flgs []cli.Flag
	for _, f := range flags {
		env := ""
		spr := strings.Split(f.GetName(), ",")
		env = prefix + "_" + strings.ToUpper(strings.Replace(spr[0], "-", "_", -1))
		switch v := f.(type) {
		case cli.IntFlag:
			flgs = append(flgs, cli.IntFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		case cli.StringFlag:
			flgs = append(flgs, cli.StringFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		case cli.BoolFlag:
			flgs = append(flgs, cli.BoolFlag{Name: v.Name, Usage: v.Usage, EnvVar: env})
		default:
			fmt.Println("unknown")
		}
	}
	return flgs
}

func main() {
	fmt.Println("Otsimo Watch")

	app := cli.NewApp()
	app.Name = "otsimo-watch"
	app.Version = Version
	app.Usage = "Otsimo Watch Server"
	app.Author = "Sercan Değirmenci <sercan@otsimo.com>"
	var flags []cli.Flag

	flags = []cli.Flag{
		cli.IntFlag{Name: "grpc-port", Value: watch.DefaultGrpcPort, Usage: "grpc server port"},
		cli.IntFlag{Name: "health-port", Value: watch.DefaultHealthPort, Usage: "health check port"},
		cli.StringFlag{Name: "tls-cert-file", Value: "", Usage: "the server's certificate file for TLS connection"},
		cli.StringFlag{Name: "tls-key-file", Value: "", Usage: "the server's private key file for TLS connection"},
		cli.StringFlag{Name: "client-id", Value: "", Usage: "client id"},
		cli.StringFlag{Name: "client-secret", Value: "", Usage: "client secret"},
		cli.StringFlag{Name: "discovery", Value: "https://connect.otsimo.com", Usage: "auth discovery url"},
		cli.BoolFlag{Name: "no-auth", Usage: "do not check token"},

		cli.StringFlag{Name: "redis-addr", Value: "localhost:6379", Usage: "redis address"},
		cli.StringFlag{Name: "redis-password", Value: "", Usage: "redis password"},
		cli.StringFlag{Name: "redis-master", Value: config.RedisMasterName, Usage: "redis master name"},
		cli.BoolFlag{Name: "redis-sentinel", Usage: "enable redis sentinel"},
		cli.IntFlag{Name: "redis-db", Value: 0, Usage: "redis db"},
		cli.BoolFlag{Name: "no-redis", Usage: "don't use redis"},
		cli.BoolFlag{Name: "debug, d", Usage: "enable verbose log"},
	}
	app.Flags = withEnvs("OTSIMO_WATCH", flags)
	app.Action = RunAction

	log.Infoln("running", app.Name, "version:", app.Version)
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
