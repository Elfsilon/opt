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
	options := &ga.Options{
		TargetFunc:          fun.NewHimmelblauFunc(),
		Space:               mat.NewSpace(-100, 100, -100, 100),
		Mode:                fun.Minimum,
		IterationsLim:       20,
		Population:          1000,
		MutationProbability: 0.3,
		MutationRate:        0.1,
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
