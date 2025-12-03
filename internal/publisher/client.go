package publisher

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/utils"
)

func StartClient(sensorNumber int) {
	fmt.Println("Starting client...")
	for i := 0; i < sensorNumber; i++ {
		fmt.Println("Generating sensor %i: ", i)
		x := utils.GenerateSensor()
		fmt.Println(x)
	}
}
