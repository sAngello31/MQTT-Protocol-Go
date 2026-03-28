# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the broker
go run cmd/broker/main.go [-p <publisher-port>] [-s <subscriber-port>]
# Defaults: publisher port 3000, subscriber port 8080

# Run the publisher simulator
go run cmd/publisher/main.go [-s <sensor-count>]
# Default: 2 sensors

# Run tests
go test ./...

# Run a single test
go test ./internal/... -run TestName

# Build all binaries
go build ./cmd/...
```

## Module

`github.com/sAngello31/MQTT-Protocol-Go`, Go 1.24.1, no external dependencies.

## Architecture

The project implements a custom MQTT broker from scratch. It has three logical components: **broker**, **publisher** (sensor simulator), and **subscriber** (TUI, not yet implemented).

### Packet Structure

MQTT packets follow the standard three-part layout defined in `internal/mqtt/models/packets/`:
- `FixedHeader` — packet type byte + flags byte + remaining length (int32)
- Variable headers — `ConnectVarHeader`, `PublishVarHeader` per packet type
- `PublisherPayload` — raw `[]byte` data

Packet type constants (e.g. `CONNECT = 0x10`) are in `packets.go`.

### Data Flow (Publisher → Broker)

1. `cmd/publisher/main.go` calls `publisher.StartClient(n)` which spawns N goroutines, one per sensor
2. Each goroutine creates a `SensorPayload` (via `utils/generator.go`), calls `EncodeBinary()`, and sends bytes into a shared channel
3. `services/traffic_controller.go:Send()` reads from the channel — **currently only prints, does not yet send to broker**
4. The broker (`internal/broker/broker.go`) starts a TCP listener via `network.StartPublisherListener(port)` and passes connections to `services/publisher.go:HandlePublisher()`
5. The handler reads raw bytes (1024-byte buffer) and prints them — **MQTT packet parsing is not yet implemented**

### Binary Encoding

`SensorPayload.EncodeBinary()` uses length-prefixed strings (uint16 big-endian + bytes) for each string field, followed by float64 (IEEE 754 big-endian) and int64 for the timestamp. This is the format the decoder must parse.

### Key Interfaces

`internal/publisher/interfaces/payload.go` defines the `Payload` interface:
```go
type Payload interface {
    ToJSON() ([]byte, error)
    EncodeBinary() ([]byte, error)
}
```

All payload types must implement this.

### Unimplemented Stubs

These files exist but are empty — implement here when extending the project:
- `internal/mqtt/services/encoder.go` — MQTT packet encoder
- `internal/mqtt/services/decoder.go` — MQTT packet decoder
- `internal/mqtt/services/subscriber/subcribe.go` — subscription logic
- `internal/publisher/services/conn.go` — publisher TCP connection service
- `cmd/suscriber/main.go` — subscriber entry point
- `cmd/gateway/main.go` — gateway entry point

### Concurrency Model

Both broker and publisher use `context.Context` with `context.WithCancel` for shutdown. OS signals (SIGINT/SIGTERM) are caught in `utils/os_signal.go` helpers and trigger the cancel function. Goroutines check `ctx.Done()` or select on the context channel to exit cleanly.

### Network Ports

Configured via CLI flags (`internal/common/flags/flags.go`):
- Publisher listener: `-p` flag, default `3000`
- Subscriber listener: `-s` flag, default `8080`

`internal/network/client.go` currently hardcodes `localhost:3000` and `localhost:8080` — these should be made configurable when implementing real publisher→broker communication.
