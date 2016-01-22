package watch

import (
	"errors"
	"io"

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

func (con *connection) receiveLoop() error {
	for {
		req, err := con.stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch req.Type {
		case apipb.WatchRequest_CREATE:
			con.send <- &apipb.WatchResponse{
				Created: true,
			}
		case apipb.WatchRequest_CANCEL:
			con.send <- &apipb.WatchResponse{
				Canceled: true,
			}
		default:
			panic("not implemented")
		}
	}
}

func (con *connection) sendLoop() {
	for {
		select {
		case c, ok := <-con.send:
			logrus.Debugln("watch_server.go: sending:", c)
			if !ok {
				return
			}
			if err := con.stream.Send(c); err != nil {
				return
			}
		case <-con.closec:
			logrus.Debugln("watch_server.go: con.closec run")
			for {
				_, ok := <-con.send
				if !ok {
					return
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
	if err != nil {
		logrus.Errorf("watch_server.go: failed to get jwt %+v", err)
		return nil, errors.New("failed to get jwt")
	}
	id, _, err := authToken(w.server.Oidc, jwt, false)
	if err != nil {
		logrus.Errorf("watch_server.go: failed to authorize user %+v", err)
		return nil, errors.New("unauthorized user")
	}
	if id == in.ProfileId {
		h.broadcast <- in
	}
	return &apipb.EmitResponse{}, nil
}

func (w *watchGrpcServer) Watch(stream apipb.WatchService_WatchServer) error {
	jwt, err := getJWTToken(stream.Context())
	if err != nil {
		logrus.Errorf("watch_server.go: failed to get jwt %+v", err)
		return errors.New("failed to get jwt")
	}
	id, _, err := authToken(w.server.Oidc, jwt, false)
	if err != nil {
		logrus.Errorf("watch_server.go: failed to authorize user %+v", err)
		return errors.New("unauthorized user")
	}
	con := &connection{
		id:     id,
		stream: stream,
		send:   make(chan *apipb.WatchResponse, sendStreamBufLen),
		closec: make(chan struct{}),
	}
	defer con.close()
	go con.sendLoop()
	return con.receiveLoop()
}
