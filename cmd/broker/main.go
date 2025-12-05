package main

import (
	"github.com/sAngello31/MQTT-Protocol-Go/internal/broker"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/common/flags"
)

func main() {
	BrokerFlags := flags.ParseBrokerFlags()
	broker.StartBroker(BrokerFlags.PublisherPort)
}
