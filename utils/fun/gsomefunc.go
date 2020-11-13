package fun

import (
	"github.com/Elfsilon/opt/utils/mat"
)

// NewSomeFunc ...
func NewSomeFunc() *Func {
	return &Func{}
}

// Func ...
type Func struct{}

// Eval ...
func (f *Func) Eval(v mat.Vec2) float64 {
	return 2*v.X*v.X + v.X*v.Y + v.Y*v.Y
}

// EvalGrad ...
func (f *Func) EvalGrad(v mat.Vec2) mat.Vec2 {
	return mat.NewVec2(4*v.X+v.Y, 2*v.Y+v.X)
}
