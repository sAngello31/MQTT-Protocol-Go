package broker

import (
	"net"
	"sync"
)

type Router struct {
	mu            sync.RWMutex
	subscriptions map[string][]net.Conn
}

func NewRouter() *Router {
	return &Router{
		subscriptions: make(map[string][]net.Conn),
	}
}

func (r *Router) Subscribe(topic string, conn net.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.subscriptions[topic] = append(r.subscriptions[topic], conn)
}

// Route forwards payload to all subscribers of the given topic.
func (r *Router) Route(topic string, payload []byte) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, conn := range r.subscriptions[topic] {
		conn.Write(payload)
	}
}
