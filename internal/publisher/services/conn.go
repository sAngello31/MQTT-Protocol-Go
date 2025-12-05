package services

import (
	"log"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/network"
)

func InitConnection() {
	conn, err := network.ConnectPublisher()
	if err != nil {
		log.Printf("Error connecting to publisher: %v", err)
	}
	defer conn.Close()
}
