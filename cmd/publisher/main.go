package main

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/common/flags"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher"
)

func main() {
	pf := flags.ParsePublisherFlags()
	fmt.Printf("Publisher started with %d sensors → broker at %s\n", pf.SensorNumber, pf.BrokerAddr)
	publisher.StartClient(pf.SensorNumber, pf.BrokerAddr)
}
