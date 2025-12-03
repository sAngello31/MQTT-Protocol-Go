package services

import "fmt"

func Send(dataChannel chan []byte) {

	for packet := range dataChannel {
		fmt.Println("Sending packet: ", packet)
		fmt.Println("Packet was sent to the broker...")
	}
}
