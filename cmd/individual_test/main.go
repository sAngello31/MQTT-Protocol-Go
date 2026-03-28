package main

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/mqtt/packets"
)

func main() {
	x := packets.DISCONNECT
	fmt.Printf("%b\n", x)
}
