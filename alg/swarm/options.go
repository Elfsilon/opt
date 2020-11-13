package swarm

import (
	"errors"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	Swarmsize     int
	Space         mat.Space
	TargetFunc    fun.TargetFunc
	Mode          string
	Phi1          float64
	Phi2          float64
	IterationsLim int
}

func (opt *Options) validate() error {
	if opt.Swarmsize <= 0 {
		return errors.New("Swarmsize must be greater than 0")
	}
	if opt.Space.Xmin > opt.Space.Xmax || opt.Space.Ymin > opt.Space.Ymax {
		return errors.New("Incorrect Options.Space")
	}
	if opt.TargetFunc == nil {
		return errors.New("Options.TargetFunc must be set")
	}
	if opt.IterationsLim <= 0 {
		return errors.New("Options.IterationsLim must be greater than 0")
	}
	if !utils.Contains(opt.Mode, []string{
		fun.Minimum,
		fun.Maximum,
	}) {
		return errors.New("Incorrect Options.Mode")
	}
	return nil
}
