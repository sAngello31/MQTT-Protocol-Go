package models

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/rand/v2"
	"time"

	"github.com/sAngello31/MQTT-Protocol-Go/internal/publisher/interfaces"
)

type SensorPayload struct {
	SensorID  string  `json:"id"`
	Topic     string  `json:"topic"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp int64   `json:"timestamp"`
}

var _ interfaces.Payload = (*SensorPayload)(nil)

func NewSensor(sensorID string, unit string, topic string) *SensorPayload {
	return &SensorPayload{
		SensorID:  sensorID,
		Topic:     topic,
		Value:     0,
		Unit:      unit,
		Timestamp: time.Now().Unix(),
	}
}

// Change it to make it more scalable lmao
func (sensor *SensorPayload) GenerateValue() {
	switch sensor.Unit {
	case "C":
		sensor.Value = 18 + rand.Float64()*12 // 18–30 °C
	case "%":
		sensor.Value = 30 + rand.Float64()*60 // 30–90 %
	default:
		sensor.Value = rand.Float64() * 100 // generic fallback
	}
}

func (sensor *SensorPayload) ToJSON() ([]byte, error) {
	return json.Marshal(sensor)
}

// Using Length-Prefixing and uint16 for the length
func (sensor *SensorPayload) EncodeBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint16(len(sensor.SensorID)))
	binary.Write(buf, binary.BigEndian, []byte(sensor.SensorID))
	binary.Write(buf, binary.BigEndian, uint16(len(sensor.Topic)))
	binary.Write(buf, binary.BigEndian, []byte(sensor.Topic))
	binary.Write(buf, binary.BigEndian, sensor.Value)
	binary.Write(buf, binary.BigEndian, uint16(len(sensor.Unit)))
	binary.Write(buf, binary.BigEndian, []byte(sensor.Unit))
	binary.Write(buf, binary.BigEndian, sensor.Timestamp)
	return buf.Bytes(), nil
}
