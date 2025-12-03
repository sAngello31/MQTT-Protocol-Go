package main

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/common/flags"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher"
)

func main() {
	publisherFlags := flags.ParsePublisherFlags()
	fmt.Println("Publisher started with", publisherFlags.SensorNumber, "sensors")
	publisher.StartClient(publisherFlags.SensorNumber)
}
