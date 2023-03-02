package reliablebroadcast

import "alpha/besteffortbroadcast"

// ============================================================
// Instantiation
// ============================================================

type Impl struct {
	ownID int

	Output Output

	beb besteffortbroadcast.Input

	delivered bool
}

func New(ownID int, beb besteffortbroadcast.Input) *Impl {
	return &Impl{
		ownID: ownID,
		beb:   beb,
	}
}

// ============================================================
// Algorithm implementation
// ============================================================

func (rb *Impl) Broadcast(message []byte) {
	msg := originalMessage{
		Message: message,
		Sender:  rb.ownID,
	}

	rb.beb.Broadcast((&msg).Serialize())
}

func (rb *Impl) Deliver(message []byte, from int) {
	if !rb.delivered {
		msg := deserializeOriginalMessage(message)
		rb.Output.Deliver(msg.Message, msg.Sender)
		rb.beb.Broadcast(message)
	}
}

// ============================================================
// Internal message type
// ============================================================

type originalMessage struct {
	Message []byte
	Sender  int
}

func (om *originalMessage) Serialize() []byte {
	// ...
	return nil
}

func deserializeOriginalMessage(data []byte) *originalMessage {
	// ...
	return nil
}
