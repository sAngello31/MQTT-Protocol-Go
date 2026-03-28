package services

import (
	"fmt"
	"net"
)

func Send(conn net.Conn, dataChannel chan []byte) {
	for packet := range dataChannel {
		// TODO: wrap packet in MQTT PUBLISH frame before sending
		_, err := conn.Write(packet)
		if err != nil {
			fmt.Println("Error sending packet to broker:", err)
			return
		}
	}
}
