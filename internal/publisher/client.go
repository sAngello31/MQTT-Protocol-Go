package publisher

import (
	"fmt"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/models"
	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/utils"
)

func StartClient(sensorNumber int) {
	var sensors [2]models.SensorPayload
	fmt.Println("Starting client...")
	for i := 0; i < sensorNumber; i++ {
		fmt.Println("Generating sensor: ", i+1)
		sensors[i] = *utils.GenerateSensor()
		sensors[i].GenerateValue()
		fmt.Println(sensors[i].EncodeBinary())
	}
}
