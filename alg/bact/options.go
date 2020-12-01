package bact

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	TargetFunc      fun.TargetFunc
	Space           mat.Space
	Mode            string
	IterationLim    int
	Population      int
	HemotaxisLim    int
	DispCount       int
	DispProbability float64
	StepSize        float64
}

func (opt *Options) validate() error {
	return nil
}
