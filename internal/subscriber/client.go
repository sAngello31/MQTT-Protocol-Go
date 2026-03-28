package subscriber

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/subscriber/utils"
)

func StartClient(brokerAddr string) {
	fmt.Println("Starting subscriber...")
	ctx, cancel := context.WithCancel(context.Background())

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	go utils.TurnOffSubscriber(sigChannel, cancel)

	// TODO: connect to broker, send SUBSCRIBE packet, receive PUBLISH packets
	<-ctx.Done()
	fmt.Println("Subscriber shut down.")
}
