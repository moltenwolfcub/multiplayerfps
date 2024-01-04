package common

import "github.com/go-gl/mathgl/mgl32"

type WorldState struct {
	Volumes []Volume

	LightCol mgl32.Vec3
}
