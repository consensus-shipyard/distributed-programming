package reliablebroadcast

import "alpha/membership"

var _ ReliableBroadcast = (*FakeReliableBroadcast)(nil)

type FakeReliableBroadcast struct {
	bid    uint64
	group  membership.Group
	sender membership.Node
}

func NewFakeReliableBroadcast(bid uint64, group membership.Group, sender membership.Node) *FakeReliableBroadcast {
	return &FakeReliableBroadcast{bid, group, sender}
}

func (frb *FakeReliableBroadcast) Send(message []byte) {
	for _, node := range frb.group.Members() {
		frb.group.Send(message, node)
	}
}

func (frb *FakeReliableBroadcast) Deliver() []byte {
	//TODO implement me
	panic("implement me")
}
