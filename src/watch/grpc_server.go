package watch

import (
	"errors"

	"github.com/otsimo/api/apipb"
	"golang.org/x/net/context"
)

type watchGrpcServer struct {
	server *Server
}

func (w *watchGrpcServer) Emit(ctx context.Context, in *apipb.EmitRequest) (*apipb.EmitResponse, error) {
	return nil, errors.New("not implemented")
}

func (w *watchGrpcServer) Watch(stream apipb.WatchService_WatchServer) error {
	return errors.New("not implemented")
}
