package main

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/common/flags"
)

func main() {
	publisherFlags := flags.ParsePublisherFlags()
	fmt.Println("Este es el publisher")
	fmt.Println("Sensor Number: ", publisherFlags.SensorNumber)
}
