package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func HandlePublisher(ln *net.TCPListener, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Publisher listener is done...")
			return
		default:
			conn, err := ln.Accept()
			if err != nil {
				if ctx.Err() != nil || errors.Is(err, net.ErrClosed) {
					fmt.Println("Publisher listener shutting down...")
					return
				}
				log.Printf("Error accepting connection: %v", err)
			}
			go handlePublisherConnection(conn)
		}
	}
}

func handlePublisherConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Printf("Error reading from publisher connection: %v", err)
		}
		return
	}
	fmt.Println(string(buf[:n]))
}
