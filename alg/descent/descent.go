package descent

import (
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newGradDescent(startPoint mat.Vec2, ratio float64) *gradDescent {
	return &gradDescent{
		oldPoint: startPoint,
		ratio:    ratio,
	}
}

type gradDescent struct {
	extremum mat.Extremum
	oldPoint mat.Vec2
	newPoint mat.Vec2
	ratio    float64
}

func (g *gradDescent) computeNextPoint(f fun.GradFun) {
	gradient := f.EvalGrad(g.oldPoint)
	gradient = gradient.Mult(g.ratio)
	nextp := g.oldPoint.Sub(gradient)
	g.newPoint = nextp
	g.extremum = mat.NewExtremum(nextp, f.Eval(nextp))
}

func (g *gradDescent) updateOldPoint() {
	g.oldPoint = g.newPoint
}
