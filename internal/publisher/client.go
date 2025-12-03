package publisher

import "fmt"

func StartClient(sensorNumber int) {
	fmt.Println("Starting client...")
	for i := 0; i < sensorNumber; i++ {
		fmt.Println("Sensor: ", i)
	}
}
