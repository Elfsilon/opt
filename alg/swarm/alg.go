package swarm

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// ParticleSwarmAlg ...
type ParticleSwarmAlg struct {
	Iterations    int
	IterationsLim int
	Swarm         *swarm
	TFunc         *fun.TargetFunc
	Mode          string
	Phi1          float64
	Phi2          float64
}

// NewParticleSwarmAlg ...
func NewParticleSwarmAlg(opt *Options) (*ParticleSwarmAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())

	swarm := generateSwarm(opt.Swarmsize, opt.Space, opt.Mode)
	if opt.EntryExtremum != nil {
		swarm.globalExt = *opt.EntryExtremum
	}
	return &ParticleSwarmAlg{
		IterationsLim: opt.IterationsLim,
		TFunc:         &opt.TargetFunc,
		Iterations:    0,
		Swarm:         &swarm,
		Mode:          opt.Mode,
		Phi1:          opt.Phi1,
		Phi2:          opt.Phi2,
	}, nil
}

// Start ...
func (alg *ParticleSwarmAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	for {
		alg.Swarm.evalExtremums(alg.Mode, *alg.TFunc)
		alg.Swarm.correctParticles(alg.Phi1, alg.Phi2)

		alg.Iterations++

		if alg.Iterations >= alg.IterationsLim {
			break
		}
	}

	return alg.Swarm.globalExt, alg.String()
}

func (alg *ParticleSwarmAlg) String() string {
	algStat := fmt.Sprintf("ALGORITHM STATS\nIterations: %v\nMode: %v\n", alg.Iterations, alg.Mode) +
		fmt.Sprintf("Extremum:\n  Coords: (%-.5f, %-.5f)\n  Value: %-.5f\n\n", alg.Swarm.globalExt.Coord.X, alg.Swarm.globalExt.Coord.Y, alg.Swarm.globalExt.Value)
	swarmStat := fmt.Sprintf("SWARM STATS\nSwarmSize: %v\nSwarm:", alg.Swarm.size)
	for _, p := range alg.Swarm.swarm {
		swarmStat += fmt.Sprintf("  Particle\n    Position: (%-.5f, %-.5f)\n    ", p.position.X, p.position.Y) +
			fmt.Sprintf("Velocity: (%-.5f, %-.5f)\n    ", p.velocity.X, p.velocity.Y) +
			fmt.Sprintf("Local Extremum:\n      Coords: (%-.5f, %-.5f)\n      Value: %-.5f", p.localExt.Coord.X, p.localExt.Coord.Y, p.localExt.Value)
	}
	return algStat + swarmStat
}
