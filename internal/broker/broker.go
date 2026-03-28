package broker

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/broker/services"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/broker/utils"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/network"
)

func StartBroker(publisherPort, subscriberPort string) {
	fmt.Println("Starting broker...")
	ctx, cancel := context.WithCancel(context.Background())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go utils.TurnOffServer(signalChannel, cancel)

	// Listeners
	listeners := []*net.TCPListener{}
	pubLn, err := network.StartPublisherListener(publisherPort)
	if err != nil {
		panic(err)
	}
	listeners = append(listeners, pubLn)
	defer pubLn.Close()
	var wg sync.WaitGroup

	wg.Add(1)
	go initPublisherListener(pubLn, &wg, ctx)
	// TODO: Add subscriber listener

	<-ctx.Done()
	for _, ln := range listeners {
		ln.Close()
	}
	wg.Wait()
}

func initPublisherListener(ln *net.TCPListener, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	services.HandlePublisher(ln, ctx)
}
