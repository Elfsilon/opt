package fun

import (
	"github.com/Elfsilon/opt/utils/mat"
)

// NewRozenbrockFunc ...
func NewRozenbrockFunc() *RozenbrockFunc {
	return &RozenbrockFunc{}
}

// RozenbrockFunc ...
type RozenbrockFunc struct{}

// Eval ...
func (sph *RozenbrockFunc) Eval(v mat.Vec2) float64 {
	return 100*(v.X*v.X-v.Y)*(v.X*v.X-v.Y) + (v.X-1)*(v.X-1) + 390
}
