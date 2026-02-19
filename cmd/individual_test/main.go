package main

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/mqtt/models/packets"
)

func main() {
	x := packets.DISCONNECT
	fmt.Printf("%b\n", x)
}
