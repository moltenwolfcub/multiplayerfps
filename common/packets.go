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
	gob.Register(ServerBoundWorldStateRequest{})
	gob.Register(ServerBoundLightingUpdate{})
	gob.Register(ServerBoundAddVolume{})

	gob.Register(ClientBoundWorldStateUpdate{})
}

type ServerBoundWorldStateRequest struct {
}

type ServerBoundLightingUpdate struct {
	Color mgl32.Vec3
}
type ServerBoundAddVolume struct {
	Volume Volume
}

type ClientBoundWorldStateUpdate struct {
	State WorldState
}
