package immune

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// ArtificialImmuneNetworkAlg ...
type ArtificialImmuneNetworkAlg struct {
	TFunc              *fun.TargetFunc
	Space              *mat.Space
	Network            *network
	Mode               string
	IterationsLim      int
	Iteration          int
	Population         int
	Best               int
	Clones             int
	MutationRate       float64
	ScatterProbability float64
	ScatterCount       int
}

// NewArtificialImmuneNetworkAlg ...
func NewArtificialImmuneNetworkAlg(opt *Options) (*ArtificialImmuneNetworkAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	network := newNetwork(opt.Population, opt.Best, opt.Clones, opt.MutationRate)

	return &ArtificialImmuneNetworkAlg{
		Network:            &network,
		TFunc:              &opt.TargetFunc,
		Space:              &opt.Space,
		Mode:               opt.Mode,
		IterationsLim:      opt.IterationsLim,
		Iteration:          0,
		Population:         opt.Population,
		Best:               opt.Best,
		Clones:             opt.Clones,
		MutationRate:       opt.MutationRate,
		ScatterProbability: opt.ScatterProbability,
		ScatterCount:       opt.ScatterCount,
	}, nil
}

// Start ...
func (a *ArtificialImmuneNetworkAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	a.Network.generateAntibodies(*a.TFunc, *a.Space)

	for {
		clones := a.Network.createClones(*a.TFunc, *a.Space, a.Mode)
		a.Network.unite(clones, *a.TFunc, *a.Space, a.Mode)
		a.Network.scatter(a.ScatterProbability, a.ScatterCount, *a.TFunc, *a.Space)

		a.Iteration++

		if a.Iteration >= a.IterationsLim {
			break
		}
	}

	return a.Network.selectBestSolution(a.Mode), a.String()
}

func (a *ArtificialImmuneNetworkAlg) String() string {
	return fmt.Sprint("ALGORITHM STATS\n") +
		fmt.Sprintf("Iterations limit: %v\n", a.IterationsLim) +
		fmt.Sprintf("Current iteration: %v\n", a.Iteration) +
		fmt.Sprintf("Searching space: %v\n", *a.Space) +
		fmt.Sprintf("Mode: %v\n", a.Mode)
}
