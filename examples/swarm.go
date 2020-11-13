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
	options := &swarm.Options{
		Swarmsize:     3,
		Space:         mat.NewSpace(-10, 10, -10, 10),
		TargetFunc:    fun.NewReverseSphereFunc(),
		Mode:          fun.Maximum,
		Phi1:          0.1,
		Phi2:          0.05,
		IterationsLim: 1000,
	}
	swarmAlg, err := swarm.NewParticleSwarmAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Particle Swarm Optimization Algorithm")
		_, stat := swarmAlg.Start()
		fmt.Println(stat)
	}
}
