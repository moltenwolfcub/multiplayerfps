package common

import (
	"encoding/gob"
	"net"

	"github.com/go-gl/mathgl/mgl32"
)

type Packet interface {
}

type RecievedPacket struct {
	Packet Packet
	Sender net.Addr
}

func RegisterPackets() {
	gob.Register(ServerBoundLightingRequest{})

	gob.Register(ClientBoundLightingUpdate{})
}

type ServerBoundLightingRequest struct {
}

type ClientBoundLightingUpdate struct {
	Color mgl32.Vec3
}
