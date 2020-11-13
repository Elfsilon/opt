package descent

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	GradFunc      fun.GradFun
	Mode          string
	IterationsLim int
	StartPoint    mat.Vec2
	Epsilon       float64
	Ratio         float64
}

func (opt *Options) validate() error {
	return nil
}
