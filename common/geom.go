package common

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Volume struct {
	Min, Max mgl32.Vec3
}

func NewVolume(min, max mgl32.Vec3) Volume {
	var calcMin, calcMax mgl32.Vec3
	if min.X() <= max.X() {
		calcMin[0] = min.X()
		calcMax[0] = max.X()
	} else {
		calcMin[0] = max.X()
		calcMax[0] = min.X()
	}
	if min.Y() <= max.Y() {
		calcMin[1] = min.Y()
		calcMax[1] = max.Y()
	} else {
		calcMin[1] = max.Y()
		calcMax[1] = min.Y()
	}
	if min.Z() <= max.Z() {
		calcMin[2] = min.Z()
		calcMax[2] = max.Z()
	} else {
		calcMin[2] = max.Z()
		calcMax[2] = min.Z()
	}
	return Volume{
		Min: calcMin,
		Max: calcMax,
	}
}

func (v Volume) Dx() float32 {
	return v.Max.X() - v.Min.X()
}
func (v Volume) Dy() float32 {
	return v.Max.Y() - v.Min.Y()
}
func (v Volume) Dz() float32 {
	return v.Max.Z() - v.Min.Z()
}

func (v Volume) Size() mgl32.Vec3 {
	return mgl32.Vec3{
		v.Max.X() - v.Min.X(),
		v.Max.Y() - v.Min.Y(),
		v.Max.Z() - v.Min.Z(),
	}
}
