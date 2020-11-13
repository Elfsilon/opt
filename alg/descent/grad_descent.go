package descent

import (
	"fmt"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Stats ...
type Stats struct {
	iter int
}

// NewStats ...
func NewStats(iter int) Stats {
	return Stats{iter: iter}
}

// GetStats ...
func (s *Stats) GetStats() string {
	return fmt.Sprintf("\n::STATS::\nCount of iterations: %v\n", s.iter)
}

// GradDescent - implementation of grad descent optimization alg
func GradDescent(f fun.GradFun, x0 mat.Vec2, e, h float64, M int) (mat.Vec2, Stats) {
	var iter int = 0
	x := x0

	for {
		grad := f.EvalGrad(x)
		xNext := x.Sub((grad.Mult(h)))

		if mat.ENorm(xNext.Sub(x)) <= e {
			return xNext, NewStats(iter)
		}

		x = xNext
		iter++
	}
}
