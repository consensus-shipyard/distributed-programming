package besteffortbroadcast

import (
	"alpha/perfectlink"
)

// ============================================================
// Instantiation
// ============================================================

type Impl struct {
	Output Output

	links map[int]perfectlink.Input

	delivered bool
}

func New(links map[int]perfectlink.Input) *Impl {
	return &Impl{
		links:     links,
		delivered: false,
	}
}

// ============================================================
// Algorithm implementation
// ============================================================

func (beb *Impl) Broadcast(message []byte) {
	for _, link := range beb.links {
		link.Send(message)
	}
}

func (il *IncomingLink) Receive(message []byte) {
	beb := il.bebImpl

	if !beb.delivered {
		beb.Output.Deliver(message, il.nodeID)
		beb.delivered = true
	}
}

// ============================================================
// Input reception
// ============================================================

func (beb *Impl) IncomingLink(nodeID int) *IncomingLink {
	return &IncomingLink{
		nodeID:  nodeID,
		bebImpl: beb,
	}
}

type IncomingLink struct {
	nodeID  int
	bebImpl *Impl
}
