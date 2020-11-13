package fun

import (
	"math"

	"github.com/Elfsilon/opt/utils/mat"
)

// NewRastriginFunc ...
func NewRastriginFunc() *RastriginFunc {
	return &RastriginFunc{}
}

// RastriginFunc ...
type RastriginFunc struct{}

// Eval ...
func (h *RastriginFunc) Eval(v mat.Vec2) float64 {
	return (v.X*v.X - 10*math.Cos(2*math.Pi*v.X) + 10) +
		(v.Y*v.Y - 10*math.Cos(2*math.Pi*v.Y) + 10) - 330
}
