package network

import (
	"fmt"
	"net"
)

var (
	PUBLISHER_PORT  = ":3000"
	SUBSCRIBER_PORT = ":8080"
)

func StartPublisherListener(port string) (*net.TCPListener, error) {
	fullPort := ":" + port
	ln, err := net.Listen("tcp", fullPort)
	if err != nil {
		panic(err)
	}
	fmt.Println("Publisher listener started on: " + fullPort)
	return ln.(*net.TCPListener), nil
}

func StartSuscriberListener(port string) (*net.TCPListener, error) {
	fullPort := ":" + port
	ln, err := net.Listen("tcp", fullPort)
	if err != nil {
		panic(err)
	}
	fmt.Println("Suscriber listener started on: " + fullPort)
	return ln.(*net.TCPListener), nil
}
