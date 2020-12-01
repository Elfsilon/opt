package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/bact"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestBact ...
func TestBact() {
	options := &bact.Options{
		TargetFunc:      fun.NewRozenbrockFunc(),
		Space:           mat.NewSpace(-10, 10, -10, 10),
		Mode:            fun.Minimum,
		Population:      100,
		HemotaxisLim:    10,
		IterationLim:    400,
		DispCount:       30,
		DispProbability: 0.3,
		StepSize:        0.015,
	}
	bactAlg, err := bact.NewBacterialForagingAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Bact Algorithm")
		extremum, _ := bactAlg.Start()
		// fmt.Println(stat)
		fmt.Println(fmt.Sprintf("%-.20f %-.20f %-.20f", extremum.Coord.X, extremum.Coord.Y, extremum.Value))
	}
}
