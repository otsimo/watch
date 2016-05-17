package watch

import (
	"golang.org/x/net/context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	redis *RedisClient
}

func NewHealthServer(server *Server) *HealthServer {
	return &HealthServer{
		redis: server.Redis,
	}
}

func (s *HealthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	if s.redis != nil {
		if err := s.redis.Health(); err != nil {
			return &healthpb.HealthCheckResponse{
				Status: healthpb.HealthCheckResponse_NOT_SERVING,
			}, nil
		}
	}
	return &healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}, nil
}
