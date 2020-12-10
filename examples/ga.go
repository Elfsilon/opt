package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/ga"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestGA ...
func TestGA() {
	entryExt := mat.NewExtremum(mat.NewVec2(0.99, 0.99), 0.0000009)
	options := &ga.Options{
		TargetFunc:             fun.NewRozenbrockFunc(),
		Space:                  mat.NewSpace(-100, 100, -100, 100),
		Mode:                   fun.Minimum,
		IterationsLim:          20,
		Population:             100,
		MutationProbability:    0.3,
		MutationRate:           0.1,
		EntryExtremum:          &entryExt,
		EntryDistributionRatio: 0.001,
	}
	genAlg, err := ga.NewGeneticAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "GA Algorithm")
		extremum, _ := genAlg.Start()
		fmt.Println(fmt.Sprintf("%-.20f %-.20f %-.20f", extremum.Coord.X, extremum.Coord.Y, extremum.Value))
	}
}
