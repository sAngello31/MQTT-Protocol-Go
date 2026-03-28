package main

import (
	"github.com/sAngello31/MQTT-Protocol-Go/internal/common/flags"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/subscriber"
)

func main() {
	sf := flags.ParseSubscriberFlags()
	subscriber.StartClient(sf.BrokerAddr)
}
