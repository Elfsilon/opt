package immune

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	TargetFunc         fun.TargetFunc
	Space              mat.Space
	Mode               string
	IterationsLim      int
	Population         int
	Best               int
	Clones             int
	MutationRate       float64
	ScatterProbability float64
	ScatterCount       int
}

func (opt *Options) validate() error {
	return nil
}
