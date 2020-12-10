package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/immune"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestImmune ...
func TestImmune() {
	entryExt := mat.NewExtremum(mat.NewVec2(0.99, 0.99), 0.0000009)
	options := &immune.Options{
		TargetFunc:             fun.NewRozenbrockFunc(),
		Space:                  mat.NewSpace(-10, 10, -10, 10),
		Mode:                   fun.Minimum,
		IterationsLim:          100,
		Population:             100,
		Best:                   20,
		Clones:                 2,
		MutationRate:           0.015,
		ScatterProbability:     0.5,
		ScatterCount:           10,
		EntryExtremum:          &entryExt,
		EntryDistributionRatio: 0.001,
	}
	immuneAlg, err := immune.NewArtificialImmuneNetworkAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Artificial Immune Network Algorithm")
		extremum, _ := immuneAlg.Start()
		// fmt.Println(stat)
		fmt.Println(extremum)
	}
}
