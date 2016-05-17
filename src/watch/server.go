package watch

import (
	"net"
	"os"

	pb "github.com/otsimo/otsimopb"

	log "github.com/Sirupsen/logrus"
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

func (s *Server) ListenGRPC() {
	grpcPort := s.Config.GetGrpcPortString()
	//Listen
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("server.go: failed to listen %v for grpc", err)
	}
	var l = &log.Logger{
		Out:       os.Stdout,
		Formatter: &log.TextFormatter{FullTimestamp: true},
		Hooks:     make(log.LevelHooks),
		Level:     log.GetLevel(),
	}
	grpclog.SetLogger(l)

	var opts []grpc.ServerOption
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(s.Config.TlsCertFile, s.Config.TlsKeyFile)
		if err != nil {
			log.Fatalf("server.go: Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	watchGrpc := &watchGrpcServer{
		server: s,
	}

	pb.RegisterWatchServiceServer(grpcServer, watchGrpc)
	log.Infof("server.go: Binding %s for grpc", grpcPort)
	if !s.Config.NoRedis {
		cl, err := NewRedisClient(s.Config)
		if err != nil {
			log.Fatalf("failed to create redis client err=%v", err)
		}
		s.Redis = cl
	}
	grpc_health_v1.RegisterHealthServer(grpcServer, NewHealthServer(s))
	go h.run()
	//Serve
	log.Fatal(grpcServer.Serve(lis))
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
