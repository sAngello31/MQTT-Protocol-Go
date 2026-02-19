package packets

type ConnectVarHeader struct {
	ProtocolName  string // MQTT
	ProtocolLevel byte   // 3(?
	ConnectFlags  byte   // clean session, will flag, will qos, will retain, password flag, username flag
	KeepAlive     uint16
}

type PublishVarHeader struct {
	TopicName string
	PacketID  uint16 // Cuando QoS > 0
}
