package packets

type FixedHeader struct {
	PacketType byte
	Flags      byte
	RemLength  int32
}
