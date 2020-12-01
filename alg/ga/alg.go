package ga

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// GeneticAlg ...
type GeneticAlg struct {
	TFunc               *fun.TargetFunc
	Space               *mat.Space
	Evolution           *evolution
	Mode                string
	Population          int
	Iteration           int
	IterationsLim       int
	MutationProbability float64
	MutationRate        float64
}

// NewGeneticAlg ...
func NewGeneticAlg(opt *Options) (*GeneticAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	evol := newEvolution(opt.Population, opt.MutationRate, opt.MutationProbability)

	return &GeneticAlg{
		TFunc:               &opt.TargetFunc,
		Space:               &opt.Space,
		Evolution:           &evol,
		Mode:                opt.Mode,
		Iteration:           0,
		Population:          opt.Population,
		IterationsLim:       opt.IterationsLim,
		MutationProbability: opt.MutationProbability,
		MutationRate:        opt.MutationRate,
	}, nil
}

// Start ...
func (a *GeneticAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	a.Evolution.init(*a.TFunc, *a.Space)

	for {
		selected := a.Evolution.selection(*a.TFunc, a.Mode)
		a.Evolution.crossover(selected, *a.TFunc, a.Mode)
		a.Evolution.mutation(*a.TFunc, a.Mode)

		a.Iteration++

		if a.Iteration >= a.IterationsLim {
			break
		}
	}

	return mat.NewExtremum(a.Evolution.population[0].position, a.Evolution.population[0].value), a.String()
}

func (a *GeneticAlg) String() string {
	return fmt.Sprint("ALGORITHM STATS\n") +
		fmt.Sprintf("Iterations limit: %v\n", a.IterationsLim) +
		fmt.Sprintf("Current iteration: %v\n", a.Iteration) +
		fmt.Sprintf("Searching space: %v\n", *a.Space) +
		fmt.Sprintf("Mode: %v\n", a.Mode)
}
