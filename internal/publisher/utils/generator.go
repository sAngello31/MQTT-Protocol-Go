package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/models"
)

var (
	level1 = [3]string{"house", "office", "plant"}
	level2 = [3]string{"floor1", "floor2", "floor3"}
	level3 = [2]map[string]string{{"topic": "temperature", "unit": "C"}, {"topic": "humidity", "unit": "%"}}
	rng    = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateSensor() *models.SensorPayload {
	id := generateSensorID()
	topic, unit := generateTopic()
	return &models.SensorPayload{
		SensorID:  id,
		Topic:     topic,
		Value:     0,
		Unit:      unit,
		Timestamp: time.Now().Unix(),
	}
}

func generateSensorID() string {
	return "sensor-" + time.Now().Format("20060102150405")
}

func generateTopic() (string, string) {
	topic, unit := pickTopicUnit(level3[:])
	return fmt.Sprintf("%s/%s/%s", pick(level1[:]), pick(level2[:]), topic), unit
}

func pick(list []string) string {
	return list[rng.Intn(len(list))]
}

func pickTopicUnit(list []map[string]string) (string, string) {
	topicMap := list[rng.Intn(len(list))]
	return topicMap["topic"], topicMap["unit"]
}
