package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/bee"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestBee ...
func TestBee() {
	entryExt := mat.NewExtremum(mat.NewVec2(0.99, 0.99), 0.0000009)
	options := &bee.Options{
		TargetFunc:             fun.NewRozenbrockFunc(),
		Space:                  mat.NewSpace(-100, 100, -100, 100),
		Mode:                   fun.Minimum,
		IterationsLim:          1000,
		Elite:                  10,
		Perspect:               10,
		EliteBee:               15,
		PerspectBee:            10,
		EliteRadius:            0.1,
		PerspectRadius:         0.2,
		EntryExtremum:          &entryExt,
		EntryDistributionRatio: 0.001,
	}
	beeAlg, err := bee.NewArtificialBeeColonyAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Artificial Bee Colony Algorithm")
		extremum, stat := beeAlg.Start()
		fmt.Println(stat)
		fmt.Println(extremum)
	}
}
