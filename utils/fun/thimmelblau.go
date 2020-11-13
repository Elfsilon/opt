package fun

import (
	"github.com/Elfsilon/opt/utils/mat"
)

// NewHimmelblauFunc ...
func NewHimmelblauFunc() *HimmelblauFunc {
	return &HimmelblauFunc{}
}

// HimmelblauFunc ...
type HimmelblauFunc struct{}

// Eval ...
func (h *HimmelblauFunc) Eval(coord mat.Vec2) float64 {
	return (coord.X*coord.X+coord.Y-11)*(coord.X*coord.X+coord.Y-11) +
		(coord.X+coord.Y*coord.Y-7)*(coord.X+coord.Y*coord.Y-7)
}
