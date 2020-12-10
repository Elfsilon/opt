package bee

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	TargetFunc             fun.TargetFunc
	Space                  mat.Space
	Mode                   string
	IterationsLim          int
	Elite                  int
	Perspect               int
	EliteBee               int
	PerspectBee            int
	EliteRadius            float64
	PerspectRadius         float64
	EntryExtremum          *mat.Extremum
	EntryDistributionRatio float64
}

func (opt *Options) validate() error {
	return nil
}
