package interfaces

type Payload interface {
	ToJSON() ([]byte, error)
	EncodeBinary() ([]byte, error)
}
