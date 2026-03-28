package broker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnections(ln *net.TCPListener, router *Router, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Broker listener shutting down...")
			return
		default:
			conn, err := ln.Accept()
			if err != nil {
				if ctx.Err() != nil || errors.Is(err, net.ErrClosed) {
					fmt.Println("Broker listener shutting down...")
					return
				}
				log.Printf("Error accepting connection: %v", err)
				continue
			}
			go handleConnection(conn, router)
		}
	}
}

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Connection error from %s: %v", conn.RemoteAddr(), err)
			}
			return
		}
		// TODO: decode MQTT packet and dispatch:
		//   CONNECT    → handshake, send CONNACK
		//   PUBLISH    → router.Route(topic, payload)
		//   SUBSCRIBE  → router.Subscribe(topic, conn)
		//   PINGREQ    → send PINGRESP
		//   DISCONNECT → close connection
		fmt.Printf("Received %d bytes from %s\n", n, conn.RemoteAddr())
		_ = router
	}
}
