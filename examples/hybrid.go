package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/hybrid"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestHybrid ...
func TestHybrid(algName1, algName2 string) {
	options := &hybrid.Options{
		TargetFunc:             fun.NewRozenbrockFunc(),
		Space:                  mat.NewSpace(-10, 10, -10, 10),
		Mode:                   fun.Minimum,
		IterationsLim:          100,
		EntryDistributionRatio: 0.001,

		SwarmSwarmsize: 100,
		SwarmPhi1:      0.1,
		SwarmPhi2:      0.05,

		GAPopulation:          100,
		GAMutationProbability: 0.3,
		GAMutationRate:        0.1,

		ImmunePopulation:         100,
		ImmuneBest:               20,
		ImmuneClones:             2,
		ImmuneMutationRate:       0.015,
		ImmuneScatterProbability: 0.5,
		ImmuneScatterCount:       10,

		BeeElite:          10,
		BeePerspect:       10,
		BeeEliteBee:       15,
		BeePerspectBee:    10,
		BeeEliteRadius:    0.1,
		BeePerspectRadius: 0.2,

		BactPopulation:      100,
		BactHemotaxisLim:    10,
		BactDispCount:       30,
		BactDispProbability: 0.3,
		BactStepSize:        0.0015,
	}
	hybridAlg, err := hybrid.NewGybridAlg(algName1, algName2, options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Artificial Immune Network Algorithm")
		extremum, _, err := hybridAlg.Start()
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(stat)
		fmt.Println(fmt.Sprintf("%-.20f %-.20f %-.20f", extremum.Coord.X, extremum.Coord.Y, extremum.Value))
	}
}
