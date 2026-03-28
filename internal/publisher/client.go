package publisher

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/services"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/utils"
)

func StartClient(sensorNumber int, brokerAddr string) {
	fmt.Println("Starting client simulation...")
	ctx, cancel := context.WithCancel(context.Background())

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	go utils.TurnOffSimulation(sigChannel, cancel)

	conn, err := services.Connect(brokerAddr)
	if err != nil {
		fmt.Println("Error connecting to broker:", err)
		return
	}
	defer conn.Close()

	dataChannel := make(chan []byte, 100)
	var wg sync.WaitGroup

	go services.Send(conn, dataChannel)
	for i := range sensorNumber {
		wg.Add(1)
		fmt.Println("Generating sensor:", i+1)
		go centralCommand(i+1, dataChannel, ctx, &wg)
	}
	wg.Wait()
	close(dataChannel)
	time.Sleep(1 * time.Second)
	fmt.Println("Client simulation is over...")
}

func centralCommand(id int, dataChannel chan []byte, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	sensor := utils.GenerateSensor()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Sensor", id, "is done...")
			return
		case <-ticker.C:
			sensor.GenerateValue()
			payload, err := sensor.EncodeBinary()
			if err != nil {
				fmt.Println("Error encoding sensor:", err)
				return
			}
			select {
			case dataChannel <- payload:
			case <-ctx.Done():
				fmt.Println("Sensor", id, "is done...")
				return
			}
		}
	}
}
