package network

import (
	"fmt"
	"net"
)

var (
	PUBLISHER_PORT  = ":3000"
	SUBSCRIBER_PORT = ":8080"
)

func StartPublisherListener() (*net.TCPListener, error) {
	ln, err := net.Listen("tcp", PUBLISHER_PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("Publisher listener started on: " + PUBLISHER_PORT)
	return ln.(*net.TCPListener), nil
}

func StartSuscriberListener() (*net.TCPListener, error) {
	ln, err := net.Listen("tcp", SUBSCRIBER_PORT)
	if err != nil {
		panic(err)
	}
	fmt.Println("Suscriber listener started on: " + SUBSCRIBER_PORT)
	return ln.(*net.TCPListener), nil
}
