package hybrid

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
			TargetFunc:             opt.TargetFunc,
			Space:                  opt.Space,
			Mode:                   opt.Mode,
			IterationsLim:          opt.IterationsLim,
			Population:             opt.GAPopulation,
			MutationProbability:    opt.GAMutationProbability,
			MutationRate:           opt.GAMutationRate,
			EntryExtremum:          opt.EntryExtremum,
			EntryDistributionRatio: opt.EntryDistributionRatio,
		})
	case "bact":
		return bact.NewBacterialForagingAlg(&bact.Options{
			TargetFunc:             opt.TargetFunc,
			Space:                  opt.Space,
			Mode:                   opt.Mode,
			IterationLim:           opt.IterationsLim,
			Population:             opt.BactPopulation,
			HemotaxisLim:           opt.BactHemotaxisLim,
			DispCount:              opt.BactDispCount,
			DispProbability:        opt.BactDispProbability,
			StepSize:               opt.BactStepSize,
			EntryExtremum:          opt.EntryExtremum,
			EntryDistributionRatio: opt.EntryDistributionRatio,
		})
	case "swarm":
		return swarm.NewParticleSwarmAlg(&swarm.Options{
			TargetFunc:             opt.TargetFunc,
			Space:                  opt.Space,
			Mode:                   opt.Mode,
			IterationsLim:          opt.IterationsLim,
			Swarmsize:              opt.SwarmSwarmsize,
			Phi1:                   opt.SwarmPhi1,
			Phi2:                   opt.SwarmPhi2,
			EntryExtremum:          opt.EntryExtremum,
			EntryDistributionRatio: opt.EntryDistributionRatio,
		})
	case "bee":
		return bee.NewArtificialBeeColonyAlg(&bee.Options{
			TargetFunc:             opt.TargetFunc,
			Space:                  opt.Space,
			Mode:                   opt.Mode,
			IterationsLim:          opt.IterationsLim,
			Elite:                  opt.BeeElite,
			Perspect:               opt.BeePerspect,
			EliteBee:               opt.BeeEliteBee,
			PerspectBee:            opt.BeePerspectBee,
			EliteRadius:            opt.BeeEliteRadius,
			PerspectRadius:         opt.BeePerspectRadius,
			EntryExtremum:          opt.EntryExtremum,
			EntryDistributionRatio: opt.EntryDistributionRatio,
		})
	case "immune":
		return immune.NewArtificialImmuneNetworkAlg(&immune.Options{
			TargetFunc:             opt.TargetFunc,
			Space:                  opt.Space,
			Mode:                   opt.Mode,
			IterationsLim:          opt.IterationsLim,
			Population:             opt.ImmunePopulation,
			Best:                   opt.ImmuneBest,
			Clones:                 opt.ImmuneClones,
			MutationRate:           opt.ImmuneMutationRate,
			ScatterProbability:     opt.ImmuneScatterProbability,
			ScatterCount:           opt.ImmuneScatterCount,
			EntryExtremum:          opt.EntryExtremum,
			EntryDistributionRatio: opt.EntryDistributionRatio,
		})
	default:
		return nil, errors.New("Cannot create algorithm")
	}
}

func startAlg(alg interface{}) (mat.Extremum, string) {
	switch alg.(type) {
	case *ga.GeneticAlg:
		algt := alg.(*ga.GeneticAlg)
		return algt.Start()
	case *bee.ArtificialBeeColonyAlg:
		algt := alg.(*bee.ArtificialBeeColonyAlg)
		return algt.Start()
	case *swarm.ParticleSwarmAlg:
		algt := alg.(*swarm.ParticleSwarmAlg)
		return algt.Start()
	case *bact.BacterialForagingAlg:
		algt := alg.(*bact.BacterialForagingAlg)
		return algt.Start()
	case *immune.ArtificialImmuneNetworkAlg:
		algt := alg.(*immune.ArtificialImmuneNetworkAlg)
		return algt.Start()
	}
	return mat.NewExtremum(mat.NewVec2(0, 0), 0), "Error"
}

// GybridAlg ...
type GybridAlg struct {
	AlgName1 string
	AlgName2 string
	Opt      *Options
}

// NewGybridAlg ...
func NewGybridAlg(algName1, algName2 string, opt *Options) (*GybridAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	return &GybridAlg{
		AlgName1: algName1,
		AlgName2: algName2,
		Opt:      opt,
	}, nil
}

// Start ...
func (a *GybridAlg) Start() (mat.Extremum, string, error) {
	rand.Seed(time.Now().UnixNano())

	alg1, err := createAlg(a.AlgName1, a.Opt)
	if err != nil {
		return mat.NewExtremum(mat.NewVec2(0, 0), 0), "", err
	}

	ext1, _ := startAlg(alg1)
	a.Opt.EntryExtremum = &ext1

	alg2, err := createAlg(a.AlgName2, a.Opt)
	if err != nil {
		return mat.NewExtremum(mat.NewVec2(0, 0), 0), "", err
	}

	ext2, stats2 := startAlg(alg2)

	return ext2, stats2, nil
}
