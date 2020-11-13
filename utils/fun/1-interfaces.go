package fun

import (
	"github.com/Elfsilon/opt/utils/mat"
)

// TargetFunc ...
type TargetFunc interface {
	Eval(mat.Vec2) float64
}

// GradFun ...
type GradFun interface {
	Eval(mat.Vec2) float64
	EvalGrad(mat.Vec2) mat.Vec2
}
