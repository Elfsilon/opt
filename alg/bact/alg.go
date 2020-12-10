package bact

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// BacterialForagingAlg ...
type BacterialForagingAlg struct {
	TFunc           *fun.TargetFunc
	Space           *mat.Space
	BacteriaSwarm   *bacteriaSwarm
	IterationLim    int
	Mode            string
	Population      int
	Iteration       int
	HemotaxisLim    int
	DispCount       int
	DispProbability float64
	StepSize        float64
}

// NewBacterialForagingAlg ...
func NewBacterialForagingAlg(opt *Options) (*BacterialForagingAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	swarm := newBacteriaSwarm(opt.HemotaxisLim)

	if opt.EntryExtremum != nil {
		swarm.initOnExtremum(opt.Population, opt.StepSize, opt.TargetFunc, *opt.EntryExtremum, opt.EntryDistributionRatio)
	} else {
		swarm.init(opt.Population, opt.StepSize, opt.TargetFunc, opt.Space)
	}

	return &BacterialForagingAlg{
		TFunc:           &opt.TargetFunc,
		Space:           &opt.Space,
		BacteriaSwarm:   &swarm,
		IterationLim:    opt.IterationLim,
		Mode:            opt.Mode,
		Iteration:       0,
		Population:      opt.Population,
		HemotaxisLim:    opt.HemotaxisLim,
		DispCount:       opt.DispCount,
		DispProbability: opt.DispProbability,
		StepSize:        opt.StepSize,
	}, nil
}

// Start ...
func (a *BacterialForagingAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	for {
		a.BacteriaSwarm.hemotaxis(*a.TFunc, a.Mode)
		a.BacteriaSwarm.reproduction(a.Mode)

		a.Iteration++

		if a.Iteration >= a.IterationLim {
			break
		}

		a.BacteriaSwarm.dispersion(a.DispProbability, a.DispCount, *a.TFunc, *a.Space)
	}

	return mat.NewExtremum(a.BacteriaSwarm.swarm[0].position, a.BacteriaSwarm.swarm[0].value), a.String()
}

func (a *BacterialForagingAlg) String() string {
	return fmt.Sprint("ALGORITHM STATS\n") +
		fmt.Sprintf("Iterations limit: %v\n", a.IterationLim) +
		fmt.Sprintf("Current iteration: %v\n", a.Iteration) +
		fmt.Sprintf("Searching space: %v\n", *a.Space) +
		fmt.Sprintf("Mode: %v\n", a.Mode)
}
