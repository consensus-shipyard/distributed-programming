package reliablebroadcast

import "net"

// ReliableBroadcast example
type ReliableBroadcast interface {
	Send(message []byte)
	Deliver() chan []byte
}

// ReliableBroadcastFactory interface.
// The problem with this pattern is that it imposes a strict interface to
// protocol instance constructors. The factory implementation constructor
// can take any parameters though.
type ReliableBroadcastFactory interface {
	New(bid uint64, sender net.Addr) ReliableBroadcast
}

// ReliableBroadcastImplFactory implements ReliableBroadcastFactory
var _ ReliableBroadcastFactory = (*ReliableBroadcastImplFactory)(nil)

type ReliableBroadcastImplFactory struct{}

func NewReliableBroadcastImplFactory(
// can take any parameter which can then be passed to the protocol instance
// constructor.
) *ReliableBroadcastImplFactory {
	return &ReliableBroadcastImplFactory{}
}

func (f *ReliableBroadcastImplFactory) New(bid uint64, sender net.Addr) ReliableBroadcast {
	return NewReliableBroadcastImpl(bid, sender)
}

// ReliableBroadcastImpl implements Reliable Broadcast
var _ ReliableBroadcast = (*ReliableBroadcastImpl)(nil)

type ReliableBroadcastImpl struct{}

func NewReliableBroadcastImpl(bid uint64, sender net.Addr) *ReliableBroadcastImpl {
	return &ReliableBroadcastImpl{}
}

func (rb *ReliableBroadcastImpl) Send(message []byte) {}
func (rb *ReliableBroadcastImpl) Deliver() []byte     { return nil }
