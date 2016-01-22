package watch

import (
	"github.com/Sirupsen/logrus"
	"github.com/otsimo/api/apipb"
)

type hub struct {
	connections map[string]*connection
	broadcast   chan *apipb.EmitRequest
	register    chan *connection
	unregister  chan *connection
}

var h = hub{
	broadcast:   make(chan *apipb.EmitRequest),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[string]*connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c.id] = c
		case c := <-h.unregister:
			if _, ok := h.connections[c.id]; ok {
				delete(h.connections, c.id)
				close(c.send)
			}
		case m := <-h.broadcast:
			if con, ok := h.connections[m.ProfileId]; ok {
				select {
				case con.send <- &apipb.WatchResponse{Event: m.Event}:
				default:
					logrus.Debugln("default broadcast select")
					go con.close()
				}
			}
		}
	}
}
