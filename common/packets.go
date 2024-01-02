package common

import (
	"net"
)

type Packet interface {
}

type RecievedPacket struct {
	Packet Packet
	Sender net.Addr
}

func RegisterPackets() {
	// gob.Register()
}
