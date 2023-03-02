package perfectlink

type Input interface {
	Send(message []byte)
}

type Output interface {
	Receive(message []byte)
}
