package membership

import "net"

// Scaffolding interfaces
type Node interface {
	Id() []byte
	NetworkAddress() net.Addr
}

type Group interface {
	Add(node Node)
	Remove(nodeId []byte)
	Members() []Node
}
