package packets

type ConnectVarHeader struct {
	ProtocolName  string
	ProtocolLevel byte
	ConnectFlags  byte
	KeepAlive     uint16
}
