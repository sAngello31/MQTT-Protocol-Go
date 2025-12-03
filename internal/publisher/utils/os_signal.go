package utils

import (
	"context"
	"fmt"
	"os"
)

func TurnOffSimulation(signalChannel chan os.Signal, cancel context.CancelFunc) {
	<-signalChannel
	fmt.Println("Killing simulation...")
	cancel()
}
