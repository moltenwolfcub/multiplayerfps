package client

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/moltenwolfcub/gogl-utils"
)

func Cuboid(size mgl32.Vec3) gogl.Object {
	o := gogl.Object{
		Type: "cuboid",
	}

	x, y, z := size.X(), size.Y(), size.Z()
	o.Verticies = []float32{
		// Back
		0, 0, 0, 0, 0,
		0, y, 0, 0, 1,
		x, y, 0, 1, 1,
		0, 0, 0, 0, 0,
		x, y, 0, 1, 1,
		x, 0, 0, 1, 0,

		// Front
		0, 0, z, 0, 0,
		x, y, z, 1, 1,
		0, y, z, 0, 1,
		0, 0, z, 0, 0,
		x, 0, z, 1, 0,
		x, y, z, 1, 1,

		// Bottom
		0, 0, 0, 0, 1,
		x, 0, z, 1, 0,
		0, 0, z, 0, 0,
		0, 0, 0, 0, 1,
		x, 0, 0, 1, 1,
		x, 0, z, 1, 0,

		// Top
		0, y, 0, 0, 1,
		0, y, z, 0, 0,
		x, y, z, 1, 0,
		0, y, 0, 0, 1,
		x, y, z, 1, 0,
		x, y, 0, 1, 1,

		// Left
		0, 0, 0, 0, 0,
		0, y, z, 1, 1,
		0, y, 0, 0, 1,
		0, 0, 0, 0, 0,
		0, 0, z, 1, 0,
		0, y, z, 1, 1,

		// Front
		x, 0, 0, 0, 0,
		x, y, 0, 0, 1,
		x, y, z, 1, 1,
		x, 0, 0, 0, 0,
		x, y, z, 1, 1,
		x, 0, z, 1, 0,
	}
	o.VertexStride = 5

	o.CalcNormals(12)
	o.FillBuffers()

	return o
}
