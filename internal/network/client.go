package network

import "net"

var (
	PUBLISHER_URL  = "localhost" + PUBLISHER_PORT
	SUBSCRIBER_URL = "localhost" + SUBSCRIBER_PORT
)

func ConnectPublisher() (*net.TCPConn, error) {
	conn, err := net.Dial("tcp", PUBLISHER_URL)
	if err != nil {
		panic(err)
	}
	return conn.(*net.TCPConn), nil
}

func ConnectSuscriber() (*net.TCPConn, error) {
	conn, err := net.Dial("tcp", SUBSCRIBER_URL)
	if err != nil {
		panic(err)
	}
	return conn.(*net.TCPConn), nil
}
