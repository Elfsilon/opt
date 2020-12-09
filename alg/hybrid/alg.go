package gybrid

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/mat"

	"github.com/Elfsilon/opt/alg/bact"
	"github.com/Elfsilon/opt/alg/bee"
	"github.com/Elfsilon/opt/alg/ga"
	"github.com/Elfsilon/opt/alg/immune"
	"github.com/Elfsilon/opt/alg/swarm"
)

func createAlg(algName string, opt *Options) (interface{}, error) {
	switch algName {
	case "ga":
		return ga.NewGeneticAlg(&ga.Options{
			TargetFunc:          opt.TargetFunc,
			Space:               opt.Space,
			Mode:                opt.Mode,
			IterationsLim:       opt.IterationsLim,
			Population:          opt.Population,
			MutationProbability: opt.MutationProbability,
			MutationRate:        opt.MutationRate,
		})
	case "bact":
		return bact.NewBacterialForagingAlg(&bact.Options{})
	case "swarm":
		return swarm.NewParticleSwarmAlg(&swarm.Options{})
	case "bee":
		return bee.NewArtificialBeeColonyAlg(&bee.Options{})
	case "immune":
		return immune.NewArtificialImmuneNetworkAlg(&immune.Options{})
	default:
		return nil, errors.New("Cannot create algorithm")
	}
}

func startAlg(alg interface{}) (mat.Extremum, string) {
	switch alg.(type) {
	case ga.GeneticAlg:
		algt := alg.(ga.GeneticAlg)
		return algt.Start()
	case bee.ArtificialBeeColonyAlg:
		algt := alg.(bee.ArtificialBeeColonyAlg)
		return algt.Start()
	case swarm.ParticleSwarmAlg:
		algt := alg.(swarm.ParticleSwarmAlg)
		return algt.Start()
	case bact.BacterialForagingAlg:
		algt := alg.(bact.BacterialForagingAlg)
		return algt.Start()
	case immune.ArtificialImmuneNetworkAlg:
		algt := alg.(immune.ArtificialImmuneNetworkAlg)
		return algt.Start()
	}
	return mat.NewExtremum(mat.NewVec2(0, 0), 0), "Error"
}

// GybridAlg ...
type GybridAlg struct {
	Alg1 interface{}
	Alg2 interface{}
}

// NewGybridAlg ...
func NewGybridAlg(algName1, algName2 string, opt *Options) (*GybridAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	alg1, err := createAlg(algName1, opt)
	if err != nil {
		return nil, err
	}

	alg2, err := createAlg(algName2, opt)
	if err != nil {
		return nil, err
	}

	return &GybridAlg{
		Alg1: alg1,
		Alg2: alg2,
	}, nil
}

// Start ...
func (a *GybridAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	ext1, stats1 := startAlg(a.Alg1)

	return ext1, stats1
}
