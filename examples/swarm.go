package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/swarm"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestSwarm ...
func TestSwarm() {
	entryExt := mat.NewExtremum(mat.NewVec2(0.99, 0.99), 0.0000009)
	options := &swarm.Options{
		Swarmsize:              3,
		Space:                  mat.NewSpace(-10, 10, -10, 10),
		TargetFunc:             fun.NewReverseSphereFunc(),
		Mode:                   fun.Maximum,
		Phi1:                   0.1,
		Phi2:                   0.05,
		IterationsLim:          1000,
		EntryExtremum:          &entryExt,
		EntryDistributionRatio: 0.001,
	}
	swarmAlg, err := swarm.NewParticleSwarmAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Particle Swarm Optimization Algorithm")
		extremum, _ := swarmAlg.Start()
		fmt.Println(fmt.Sprintf("%-.20f %-.20f %-.20f", extremum.Coord.X, extremum.Coord.Y, extremum.Value))
	}
}
