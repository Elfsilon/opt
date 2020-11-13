package descent

import (
	"fmt"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// NewGradDescentAlg ...
func NewGradDescentAlg(opt *Options) (*GradDescentAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	descent := newGradDescent(opt.StartPoint, opt.Ratio)

	return &GradDescentAlg{
		GradDescent:   descent,
		GFunc:         opt.GradFunc,
		StartPoint:    opt.StartPoint,
		Mode:          opt.Mode,
		IterationsLim: opt.IterationsLim,
		Iteration:     0,
		Epsilon:       opt.Epsilon,
		Ratio:         opt.Ratio,
	}, nil
}

// GradDescentAlg ...
type GradDescentAlg struct {
	GFunc         fun.GradFun
	GradDescent   *gradDescent
	StartPoint    mat.Vec2
	Mode          string
	IterationsLim int
	Iteration     int
	Epsilon       float64
	Ratio         float64
}

// Start ...
func (g *GradDescentAlg) Start() (mat.Extremum, string) {
	for {
		g.GradDescent.computeNextPoint(g.GFunc)
		if mat.ENorm(g.GradDescent.newPoint.Sub(g.GradDescent.oldPoint)) <= g.Epsilon || g.Iteration >= g.IterationsLim {
			break
		}
		g.GradDescent.updateOldPoint()
		g.Iteration++
	}

	return g.GradDescent.extremum, g.String()
}

func (g *GradDescentAlg) String() string {
	return fmt.Sprint("ALGORITHM STATS\n") +
		fmt.Sprintf("Iterations limit: %v\n", g.IterationsLim) +
		fmt.Sprintf("Current iteration: %v\n", g.Iteration) +
		fmt.Sprintf("Mode: %v\n", g.Mode) +
		fmt.Sprintf("Start point: %v\n", g.StartPoint) +
		fmt.Sprintf("Epsilon: %v\n", g.Epsilon) +
		fmt.Sprintf("Ratio: %v\n", g.Ratio) +
		fmt.Sprintf("Found extremum: %v\n", g.GradDescent.extremum)
}
