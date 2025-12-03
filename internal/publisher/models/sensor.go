package models

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"time"
)

type SensorPayload struct {
	SensorID  string  `json:"id"`
	Topic     string  `json:"topic"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Timestamp int64   `json:"timestamp"`
}

func NewSensor(sensorID string, value float64, unit string, topic string) *SensorPayload {
	return &SensorPayload{
		SensorID:  sensorID,
		Topic:     topic,
		Value:     value,
		Unit:      unit,
		Timestamp: time.Now().Unix(),
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
