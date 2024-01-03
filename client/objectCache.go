package client

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/moltenwolfcub/gogl-utils"
)

type ObjectParameters struct {
	// All parameters that any object uses will be in here.
	// Only set the ones relevant to the object being used.
	Size1 float32
	Size3 mgl32.Vec3
}

func (o ObjectParameters) String() string {
	str := "{"
	var empty ObjectParameters

	if o.Size1 != empty.Size1 {
		str += fmt.Sprintf("Size1: %v, ", o.Size1)
	}
	if o.Size3 != empty.Size3 {
		str += fmt.Sprintf("Size3: %v, ", o.Size3)
	}

	str += "}"

	return str
}

type ObjectKey struct {
	ObjectType string
	Params     ObjectParameters
}

func (o ObjectKey) String() string {
	return fmt.Sprintf("{%s:%v}", o.ObjectType, o.Params)
}

type ObjectCache map[ObjectKey]*gogl.Object

func (c ObjectCache) GetOrCreate(key ObjectKey) (*gogl.Object, error) {
	obj, ok := c[key]
	if ok {
		return obj, nil
	}
	switch key.ObjectType {
	case "cube":
		o := gogl.Cube(key.Params.Size1)
		c[key] = &o
		return &o, nil

	case "pentahedron":
		o := gogl.Pentahedron(key.Params.Size1)
		c[key] = &o
		return &o, nil

	case "cuboid":
		o := Cuboid(key.Params.Size3)
		c[key] = &o
		return &o, nil

	default:
		return &gogl.Object{}, fmt.Errorf("unknown ObjectType(%s) in key", key.ObjectType)
	}
}

func (c ObjectCache) String() string {
	str := "objectCache["

	for k := range c {
		str += fmt.Sprintf("%v", k)
	}
	return str
}
