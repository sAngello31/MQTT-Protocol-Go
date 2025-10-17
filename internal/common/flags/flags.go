package flags

import (
	"flag"
	"os"
)

type PublisherFlags struct {
	SensorNumber int
}

func ParsePublisherFlags() *PublisherFlags {
	description := "Number of sensors that the publisher will generate"
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	csf := &PublisherFlags{}
	fs.IntVar(&csf.SensorNumber, "sensors", 2, description)
	fs.Parse(os.Args[1:])
	return csf
}
