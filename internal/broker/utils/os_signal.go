package utils

import (
	"context"
	"fmt"
	"os"
)

func TurnOffServer(signalChannel chan os.Signal, cancel context.CancelFunc) {
	<-signalChannel
	fmt.Println("Killing Broker Server...")
	cancel()
}
