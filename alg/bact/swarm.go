package bact

import (
	"math/rand"
	"sort"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newBacteriaSwarm(hemotaxisLim int) bacteriaSwarm {
	return bacteriaSwarm{
		swarm:        make([]bacterium, 0),
		hemotaxisLim: hemotaxisLim,
	}
}

type bacteriaSwarm struct {
	swarm        []bacterium
	hemotaxisLim int
}

func (b *bacteriaSwarm) init(population int, stepsize float64, f fun.TargetFunc, s mat.Space) {
	for i := 0; i < population; i++ {
		bact := newBacterium()
		bact.init(stepsize, f, s)
		b.swarm = append(b.swarm, bact)
	}
}

func (b *bacteriaSwarm) hemotaxis(f fun.TargetFunc, mode string) {
	for i := range b.swarm {
		b.swarm[i].hemotaxis(b.hemotaxisLim, f, mode)
	}
}

func (b *bacteriaSwarm) reproduction(mode string) {
	if mode == fun.Minimum {
		sort.SliceStable(b.swarm, func(i, j int) bool {
			return b.swarm[i].value < b.swarm[j].value
		})
	} else {
		sort.SliceStable(b.swarm, func(i, j int) bool {
			return b.swarm[i].value > b.swarm[j].value
		})
	}

	bestHalf := b.swarm[:len(b.swarm)/2]
	newPopulation := make([]bacterium, 0)

	for i := 0; i < len(bestHalf); i++ {
		newPopulation = append(newPopulation, bestHalf[i], bestHalf[i])
	}

	b.swarm = newPopulation
}

func (b *bacteriaSwarm) dispersion(probability float64, quantity int, f fun.TargetFunc, s mat.Space) {
	for i := 0; i < quantity; i++ {
		p := utils.RandFloat(0, 1)
		if p >= probability {
			index := rand.Intn(len(b.swarm))
			b.swarm[index].randPosition(f, s)
		}
	}
}
