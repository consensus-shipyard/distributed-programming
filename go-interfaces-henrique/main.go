package main

import (
	"alpha/membership"
	"alpha/reliablebroadcast"
)

func SomeUsage(rbFactory reliablebroadcast.ReliableBroadcastFactory, numInstances uint) {
	for i := 0; i < int(numInstances); i++ {
		bid := uint64(i)
		rb := rbFactory.New(bid, nil)
		rb.Deliver()
	}
}

func main() {

}

// ConsensusImpl implements Consensus interface
type ConsensusImpl struct {
	group   membership.Group
	factory reliablebroadcast.ReliableBroadcastFactory
}

func NewConsensusImpl(factory reliablebroadcast.ReliableBroadcastFactory, group membership.Group) *ConsensusImpl {
	return &ConsensusImpl{}
}

func (c *ConsensusImpl) Propose(value []byte) {
	rb.Send(value)

}

func (c *ConsensusImpl) Deliver() {
	futures := make([]chan []byte)
	for i, node := range c.group.Members() {
		rb := c.factory.New(i, node.NetworkAddress())
		if i == MyID() {
			rb.Send(value)
		}
		future := rb.Deliver()
	}
}