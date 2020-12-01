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
	options := &immune.Options{
		TargetFunc:    fun.NewReverseSphereFunc(),
		Space:         mat.NewSpace(-10, 10, -10, 10),
		Mode:          fun.Maximum,
		IterationsLim: 10000,
		Population:    200,
		Best:          10,
		Clones:        20,
		MutationRate:  0.0025,
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
