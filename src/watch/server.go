package watch

import (
	"fmt"
	"net"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/otsimo/health"
	pb "github.com/otsimo/otsimopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	Config *Config
	Oidc   *Client
	Redis  *RedisClient
	NoAuth bool
}

func init() {
	var l = &log.Logger{
		Out:       os.Stdout,
		Formatter: &log.TextFormatter{FullTimestamp: true},
		Hooks:     make(log.LevelHooks),
		Level:     log.GetLevel(),
	}
	grpclog.SetLogger(l)
}

func (s *Server) Healthy() error {
	return nil
}

func (s *Server) ListenGRPC() error {
	grpcPort := s.Config.GetGrpcPortString()
	//Listen
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("server.go: failed to listen, %v", err)
	}
	hs := health.New()
	var opts []grpc.ServerOption
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(s.Config.TlsCertFile, s.Config.TlsKeyFile)
		if err != nil {
			return fmt.Errorf("server.go: Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	watchGrpc := &watchGrpcServer{
		server: s,
	}

	pb.RegisterWatchServiceServer(grpcServer, watchGrpc)
	grpc_health_v1.RegisterHealthServer(grpcServer, hs)

	if !s.Config.NoRedis {
		cl, err := NewRedisClient(s.Config)
		if err != nil {
			return fmt.Errorf("failed to create redis client, %v", err)
		}
		s.Redis = cl
		hs.Checks = append(hs.Checks, s.Redis)
	}

	log.Infof("server.go: Binding %s for grpc and :%d for health", grpcPort, s.Config.HealthPort)
	go http.ListenAndServe(s.Config.GetHealthPortString(), hs)
	go h.run()
	return grpcServer.Serve(lis)
}

func NewServer(config *Config) *Server {
	server := &Server{
		Config: config,
		NoAuth: config.NoAuth,
	}
	if !server.NoAuth {
		c, err := NewOIDCClient(config.ClientID, config.ClientSecret, config.AuthDiscovery)
		if err != nil {
			log.Fatal("Unable to create Oidc client", err)
		}
		server.Oidc = c
	}
	return server
}

func (s *Server) Emit(in *pb.EmitRequest) {
	if s.Config.NoRedis {
		h.broadcast <- in
	} else {
		s.Redis.Emit(in)
	}
}
