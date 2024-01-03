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
	gob.Register(ServerBoundLightingUpdate{})
	gob.Register(ServerBoundWorldStateRequest{})

	gob.Register(ClientBoundLightingUpdate{})
	gob.Register(ClientBoundWorldStateUpdate{})
}

type ServerBoundLightingRequest struct {
}
type ServerBoundWorldStateRequest struct {
}

type ServerBoundLightingUpdate struct {
	Color mgl32.Vec3
}

type ClientBoundLightingUpdate struct {
	Color mgl32.Vec3
}
type ClientBoundWorldStateUpdate struct {
	State WorldState
}
