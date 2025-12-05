package flags

import (
	"flag"
	"os"
)

type PublisherFlags struct {
	SensorNumber int
}

type BrokerFlags struct {
	PublisherPort  string
	SubscriberPort string
}

func ParsePublisherFlags() *PublisherFlags {
	description := "Number of sensors that the publisher will generate"
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	csf := &PublisherFlags{}
	fs.IntVar(&csf.SensorNumber, "s", 2, description)
	fs.Parse(os.Args[1:])
	return csf
}

func ParseBrokerFlags() *BrokerFlags {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	bf := &BrokerFlags{}
	fs.StringVar(&bf.PublisherPort, "p", "3000", "Publisher port of the broker")
	fs.StringVar(&bf.SubscriberPort, "s", "8080", "Subscriber port of the broker")
	fs.Parse(os.Args[1:])
	return bf
}
