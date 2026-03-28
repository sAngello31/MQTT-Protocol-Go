package utils

import (
	"context"
	"fmt"
	"os"
)

func TurnOffSubscriber(signalChannel chan os.Signal, cancel context.CancelFunc) {
	<-signalChannel
	fmt.Println("Shutting down subscriber...")
	cancel()
}
