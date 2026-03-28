package broker

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/broker/utils"
)

func StartBroker(port string) {
	fmt.Println("Starting broker...")
	ctx, cancel := context.WithCancel(context.Background())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go utils.TurnOffServer(signalChannel, cancel)

	router := NewRouter()

	addr := ":" + port
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Broker listening on", addr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleConnections(ln.(*net.TCPListener), router, ctx)
	}()

	<-ctx.Done()
	ln.Close()
	wg.Wait()
	fmt.Println("Broker shut down.")
}
