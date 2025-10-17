# MQTT-Protocol-Go
- Version 1.0

## Stack y Librerias
- Go 1.22.2
- Bubble Library para la TUI del cliente

## Requerimientos Funcionales
### Generales
1. Generar Logs

### Client/Subscriber
1. TUI para la interaccion con el usuario.
2. Conexión al broker.
3. Suscripción a topics.
4. Publicación de mensajes.
4.1. A tiempo real.
4.2. Permitir un pipe para el guardado de mensajes en un archivo (Futuro).

Por el momento unicamente tener un TUI que permita al usuario ver e interactuar con el broker

### Broker
1. Leer el standar Oasis
1.1. CONNECT/ CONNACK
1.2. PUBLISH/ PUBACK
1.3. SUBSCRIBE/ SUBACK
1.4. UNSUBSCRIBE/ UNSUBACK
1.5. DISCONNECT

2. Calidad de Servicio (QoS)
2.1. QoS 0: Fire and Forget
2.2. QoS 1: PUBACK
2.3. QoS 2: PUBREC, PUBREL, PUBCOMP - Para luego puesto que esto es medio complicado implementarlo.

3. Tabla de Suscripciones.
4. Retained Messages - Para cuando un user se desconecta y tiene la flag retained, el broker mantiene el mensaje.
5. Session Persistence (Si se implementa Qos > 0)

### Gateway
1. Conexión al broker.
2. Suscripción a topics.
3. Publicación de mensajes.
4. Posiblemente agregar metadata (Esta por verse)

### Publisher
1. Generar entre 1 a 2 sensores simulados (Temperatura y Humedad).
2. Trabajar en dos hilos distintos.
3. Conectarse con el Gateway.
4. Mandar mensajes a los topics correspondientes.
5. Por el momento se debe generar el topic de los sensores de forma aleatoria.
6. Por default se crearan dos sensores (Temperatura y Humedad). Pero el usuario puede indicar cuantos sensores desea generar con flags.

### Arquitectura del proyecto
mqtt-system/
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
│   │   ├── server.go        # Lógica del broker MQTT
│   │   ├── client_manager.go
│   │   └── message_store.go
│   │
│   ├── gateway/
│   │   ├── load_balancer.go # Reenvía mensajes a brokers
│   │   ├── routing.go
│   │   └── metrics.go
│   │
│   ├── publisher/
│   │   ├── client.go        # Conecta con gateway o broker
│   │   └── message.go
│   │
│   ├── subscriber/
│   │   ├── client.go
│   │   └── handler.go
│   │
│   ├── mqtt/
│   │   ├── protocol.go      # Parseo, encoding/decoding de paquetes MQTT
│   │   └── utils.go
│   │
│   ├── network/
│   │   ├── connection.go    # Conexión TCP y manejo de sockets
│   │   └── pool.go          # Reuso de conexiones si aplica
│   │
│   └── common/
│       ├── config.go        # Carga de configuración (YAML, JSON, ENV)
│       ├── logger.go        # Logging centralizado
│       └── constants.go
│
├── pkg/                     # Paquetes reutilizables opcionales
│   ├── metrics/
│   └── utils/
│
├── go.mod
└── go.sum