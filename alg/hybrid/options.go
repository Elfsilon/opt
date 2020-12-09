package gybrid

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	TargetFunc          fun.TargetFunc
	Space               mat.Space
	Mode                string
	IterationsLim       int
	Population          int
	MutationProbability float64
	MutationRate        float64
}

func (opt *Options) validate() error {
	return nil
}
