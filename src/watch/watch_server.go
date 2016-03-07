package watch

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
	"golang.org/x/net/context"
)

const (
	sendStreamBufLen = 16
)

type connection struct {
	id     string
	stream apipb.WatchService_WatchServer
	send   chan *apipb.WatchResponse
	closec chan struct{}
}

func (con *connection) close() {
	close(con.send)
	close(con.closec)
	h.unregister <- con
}

func (con *connection) sendLoop() error {
	for {
		select {
		case c, ok := <-con.send:
			logrus.Debugf("watch_server.go: sending: %v", c)
			if !ok {
				return errors.New("internal")
			}
			if err := con.stream.Send(c); err != nil {
				return err
			}
		case <-con.closec:
			logrus.Debugln("watch_server.go: con.closec run")
			for {
				_, ok := <-con.send
				if !ok {
					return errors.New("close error")
				}
			}
		}
	}
}

type watchGrpcServer struct {
	server *Server
}

func (w *watchGrpcServer) Emit(ctx context.Context, in *apipb.EmitRequest) (*apipb.EmitResponse, error) {
	jwt, err := getJWTToken(ctx)
	logrus.Debugf("watch_server.go: Emit %+v", in)
	if err != nil {
		logrus.Errorf("watch_server.go: failed to get jwt %+v", err)
		return nil, errors.New("failed to get jwt")
	}
	_, _, err = authToken(w.server.Oidc, jwt, false)
	if err != nil {
		logrus.Errorf("watch_server.go: failed to authorize user %+v", err)
		return nil, errors.New("unauthorized user")
	}

	h.broadcast <- in
	return &apipb.EmitResponse{}, nil
}

func (w *watchGrpcServer) Watch(req *apipb.WatchRequest, stream apipb.WatchService_WatchServer) error {
	jwt, err := getJWTToken(stream.Context())
	if err != nil {
		logrus.Errorf("watch_server.go: failed to get jwt %+v", err)
		return errors.New("failed to get jwt")
	}
	id, _, err := authToken(w.server.Oidc, jwt, false)
	if err != nil && !w.server.NoAuth {
		logrus.Errorf("watch_server.go: failed to authorize user %+v", err)
		return errors.New("unauthorized user")
	}

	con := &connection{
		id:     id,
		stream: stream,
		send:   make(chan *apipb.WatchResponse, sendStreamBufLen),
		closec: make(chan struct{}),
	}

	logrus.Debugf("watch_server.go: started to watch %s", id)
	h.register <- con

	defer con.close()
	return con.sendLoop()
}
