package main

import (
	"fmt"
	"os"
	"watch"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Version string
var config *watch.Config = watch.NewConfig()

const (
	EnvDebugName    = "OTSIMO_WATCH_DEBUG"
	EnvGrpcPortName = "OTSIMO_WATCH_GRPC_PORT"
)

func RunAction(c *cli.Context) {
	config.Debug = c.Bool("debug")
	config.GrpcPort = c.Int("grpc-port")
	config.TlsCertFile = c.String("tls-cert-file")
	config.TlsKeyFile = c.String("tls-key-file")

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	server := watch.NewServer(config)
	server.ListenGRPC()
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
		cli.IntFlag{Name: "grpc-port", Value: watch.DefaultGrpcPort, Usage: "grpc server port", EnvVar: EnvGrpcPortName},
		cli.StringFlag{Name: "tls-cert-file", Value: "", Usage: "the server's certificate file for TLS connection"},
		cli.StringFlag{Name: "tls-key-file", Value: "", Usage: "the server's private key file for TLS connection"},
		cli.BoolFlag{Name: "debug, d", Usage: "enable verbose log", EnvVar: EnvDebugName},
	}
	app.Flags = flags
	app.Action = RunAction
	app.Run(os.Args)
}

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
