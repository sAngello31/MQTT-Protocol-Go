package flags

import "flag"

type PublisherFlags struct {
	SensorNumber int
}

type BrokerFlags struct {
	PublisherPort  string
	SubscriberPort string
}

func ParsePublisherFlags() *PublisherFlags {
	csf := &PublisherFlags{}
	flag.IntVar(&csf.SensorNumber, "s", 2, "Number of sensors that the publisher will generate")
	flag.Parse()
	return csf
}

func ParseBrokerFlags() *BrokerFlags {
	bf := &BrokerFlags{}
	flag.StringVar(&bf.PublisherPort, "p", "3000", "Publisher port of the broker")
	flag.StringVar(&bf.SubscriberPort, "sub", "8080", "Subscriber port of the broker")
	flag.Parse()
	return bf
}
