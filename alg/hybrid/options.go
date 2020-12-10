package hybrid

import (
	// "errors"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// Options ...
type Options struct {
	TargetFunc               fun.TargetFunc
	Space                    mat.Space
	Mode                     string
	IterationsLim            int
	SwarmSwarmsize           int
	SwarmPhi1                float64
	SwarmPhi2                float64
	GAPopulation             int
	GAMutationProbability    float64
	GAMutationRate           float64
	ImmunePopulation         int
	ImmuneBest               int
	ImmuneClones             int
	ImmuneMutationRate       float64
	ImmuneScatterProbability float64
	ImmuneScatterCount       int
	BeeElite                 int
	BeePerspect              int
	BeeEliteBee              int
	BeePerspectBee           int
	BeeEliteRadius           float64
	BeePerspectRadius        float64
	BactPopulation           int
	BactHemotaxisLim         int
	BactDispCount            int
	BactDispProbability      float64
	BactStepSize             float64
	EntryExtremum            *mat.Extremum
	EntryDistributionRatio   float64
}

func (opt *Options) validate() error {
	return nil
}
