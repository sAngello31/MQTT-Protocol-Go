package flags

import "flag"

type BrokerFlags struct {
	Port string
}

type PublisherFlags struct {
	SensorNumber int
	BrokerAddr   string
}

type SubscriberFlags struct {
	BrokerAddr string
}

func ParseBrokerFlags() *BrokerFlags {
	bf := &BrokerFlags{}
	flag.StringVar(&bf.Port, "p", "1883", "Broker TCP port")
	flag.Parse()
	return bf
}

func ParsePublisherFlags() *PublisherFlags {
	pf := &PublisherFlags{}
	flag.IntVar(&pf.SensorNumber, "s", 2, "Number of sensors the publisher will simulate")
	flag.StringVar(&pf.BrokerAddr, "b", "localhost:1883", "Broker address (host:port)")
	flag.Parse()
	return pf
}

func ParseSubscriberFlags() *SubscriberFlags {
	sf := &SubscriberFlags{}
	flag.StringVar(&sf.BrokerAddr, "b", "localhost:1883", "Broker address (host:port)")
	flag.Parse()
	return sf
}
