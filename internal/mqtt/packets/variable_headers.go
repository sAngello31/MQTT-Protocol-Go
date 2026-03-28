package packets

type ConnectVarHeader struct {
	ProtocolName  string // MQTT
	ProtocolLevel byte   // Version (4 = MQTT 3.1.1)
	ConnectFlags  byte   // clean session, will flag, will qos, will retain, password flag, username flag
	KeepAlive     uint16
}

type PublishVarHeader struct {
	TopicName string
	PacketID  uint16 // Only used when QoS > 0
}

type SubscribeVarHeader struct {
	PacketID uint16
}
