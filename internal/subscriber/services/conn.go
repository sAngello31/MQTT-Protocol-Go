package services

import (
	"fmt"
	"net"
)

func Connect(brokerAddr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", brokerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to broker at %s: %w", brokerAddr, err)
	}
	return conn, nil
}
