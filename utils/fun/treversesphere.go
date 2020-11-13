package fun

import (
	"github.com/Elfsilon/opt/utils/mat"
)

// NewReverseSphereFunc ...
func NewReverseSphereFunc() *ReverseSphere {
	return &ReverseSphere{}
}

// ReverseSphere ...
type ReverseSphere struct{}

// Eval ...
func (sph *ReverseSphere) Eval(coord mat.Vec2) float64 {
	return -(coord.X*coord.X + coord.Y*coord.Y)
}
