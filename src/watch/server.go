package watch

import (
	"net"
	"os"

	pb "github.com/otsimo/api/apipb"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	Config *Config
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
	//Serve
	grpcServer.Serve(lis)
}

func NewServer(config *Config) *Server {
	server := &Server{
		Config: config,
	}
	return server
}
