package besteffortbroadcast

type Input interface {
	Broadcast(message []byte)
}

type Output interface {
	Deliver(message []byte, from int)
}
