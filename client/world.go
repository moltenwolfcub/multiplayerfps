package client

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/moltenwolfcub/gogl-utils"
	"github.com/moltenwolfcub/multiplayerfps/common"
)

type ObjectInstance struct {
	Parent      *gogl.Object
	ModelMatrix mgl32.Mat4
}

func (o ObjectInstance) String() string {
	str := "{"

	str += fmt.Sprintf("%v at", o.Parent.Type)
	str += fmt.Sprintf("%v", o.ModelMatrix.Col(3).Vec3())

	str += "}"
	return str
}

type worldState struct {
	cache ObjectCache

	objects  []ObjectInstance
	lightCol mgl32.Vec3
}

func NewWorldState(state common.WorldState) (worldState, error) {
	w := worldState{
		cache: ObjectCache{},
	}

	err := w.Update(state)
	if err != nil {
		return worldState{}, err
	}

	return w, nil
}

func (w worldState) String() string {
	str := "worldState{"

	str += fmt.Sprintf("lightCol[%v]", w.lightCol)

	str += "objects["
	for _, obj := range w.objects {
		str += fmt.Sprintf("%v, ", obj)
	}
	str += "] "

	str += fmt.Sprintf("%v", w.cache)

	str += "}"
	return str
}

func (w *worldState) Update(state common.WorldState) error {
	w.objects = make([]ObjectInstance, 0)

	for _, vol := range state.Volumes {
		k := ObjectKey{
			ObjectType: "cuboid",
			Params: ObjectParameters{
				Size3: vol.Size(),
			},
		}
		obj, err := w.cache.GetOrCreate(k)
		if err != nil {
			return err
		}
		instance := ObjectInstance{
			Parent:      obj,
			ModelMatrix: mgl32.Ident4().Mul4(mgl32.Translate3D(vol.Min.X(), vol.Min.Y(), vol.Min.Z())),
		}
		w.objects = append(w.objects, instance)
	}

	w.lightCol = state.LightCol

	return nil
}
