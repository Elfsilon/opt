package examples

import (
	"log"
	"time"

	"github.com/Elfsilon/opt/alg/descent"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// TestGradDescent ...
func TestGradDescent() {
	options := &descent.Options{
		GradFunc:      fun.NewSomeFunc(),
		StartPoint:    mat.NewVec2(10, 12),
		Mode:          fun.Maximum,
		IterationsLim: 10000,
		Epsilon:       0.00001,
		Ratio:         0.001,
	}
	gdescAlg, err := descent.NewGradDescentAlg(options)
	if err != nil {
		log.Println(err)
	} else {
		defer utils.TimeTrack(time.Now(), "Particle Swarm Optimization Algorithm")
		_, stats := gdescAlg.Start()
		log.Println(stats)
	}
}
