# MQTT-Protocol-Go
- Version 1.0

***This project is currently under development.***

## Introduction
This repository contains the implementation of an MQTT (Message Queuing Telemetry Transport) protocol in Go. It includes a broker, a publisher for simulating sensor data, and a subscriber client with a Text-based User Interface (TUI). The goal is to create a lightweight and functional MQTT system for IoT applications.

## Stack and Libraries
- Go 1.22.2

## Functional Requirements
### General
1.  Generate Logs

### Client/Subscriber
1.  TUI for user interaction.
2.  Connection to the broker.
3.  Subscription to topics.
4.  Publication of messages.
    4.1. In real-time.
    4.2. Allow a pipe to save messages to a file (Future).

For now, the focus is on a TUI that allows the user to see and interact with the broker.

### Broker
1.  Adhere to the Oasis standard.
    1.1. CONNECT / CONNACK
    1.2. PUBLISH / PUBACK
    1.3. SUBSCRIBE / SUBACK
    1.4. UNSUBSCRIBE / UNSUBACK
    1.5. DISCONNECT
2.  Quality of Service (QoS)
    2.1. QoS 0: At most once (Fire and Forget).
    2.2. QoS 1: At least once (Acknowledged delivery with PUBACK).
    2.3. QoS 2: Exactly once (Assured delivery with PUBREC, PUBREL, PUBCOMP) - To be implemented later due to complexity.
3.  Subscription Table.
4.  Retained Messages: When a user with the retained flag disconnects, the broker stores the message.
5.  Session Persistence (If QoS > 0 is implemented).

### Publisher
1.  Generate 1 to 2 simulated sensors (e.g., Temperature and Humidity).
2.  Run in separate threads.
3.  Connect to the Gateway.
4.  Send messages to corresponding topics.
5.  For now, sensor topics are generated randomly.
6.  By default, two sensors (Temperature and Humidity) will be created. The user can specify the number of sensors to generate using flags.

### Project Architecture
```
mqtt-protocol-go/
├── cmd/
│   ├── broker/
│   │   └── main.go
│   ├── gateway/
│   │   └── main.go
│   ├── publisher/
│   │   └── main.go
│   └── subscriber/
│       └── main.go
│
├── internal/
│   ├── broker/
│   │   ├── broker.go
│   │   ├── services/
│   │   │   └── publisher.go
│   │   └── utils/
│   │       └── os_signal.go
│   │
│   ├── common/
│   │   └── flags/
│   │       └── flags.go
│   │
│   ├── mqtt/
│   │   └── models/
│   │       ├── mqtt_tree.go
│   │       └── node.go
│   │
│   ├── network/
│   │   ├── client.go
│   │   └── server.go
│   │
│   ├── publisher/
│   │   ├── client.go
│   │   ├── interfaces/
│   │   │   └── payload.go
│   │   ├── models/
│   │   │   ├── publisher_packet.go
│   │   │   └── sensor.go
│   │   ├── services/
│   │   │   ├── conn.go
│   │   │   └── traffic_controller.go
│   │   └── utils/
│   │       ├── generator.go
│   │       └── os_signal.go
│   │
│   └── subscriber/
│
├── pkg/
├── go.mod
└── .gitignore
```
