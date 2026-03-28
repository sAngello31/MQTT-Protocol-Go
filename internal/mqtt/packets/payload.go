package packets

type PublishPayload struct {
	Data []byte
}

type SubscribePayload struct {
	TopicFilter string
	QoS         byte
}
