package immune

import (
	"sort"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newNetwork(
	population,
	bestQuantity,
	cloneQuantity int,
	mutationRate float64,
) network {
	return network{
		population:    population,
		bestQuantity:  bestQuantity,
		cloneQuantity: cloneQuantity,
		mutationRate:  mutationRate,
	}
}

type network struct {
	antibodies    []cell
	population    int
	bestQuantity  int
	cloneQuantity int
	mutationRate  float64
}

func (n *network) generateAntibodies(f fun.TargetFunc, s mat.Space) {
	for i := 0; i < n.population; i++ {
		cell := generateCell(f, s)
		n.antibodies = append(n.antibodies, cell)
	}
}

func (n *network) createClones(f fun.TargetFunc, s mat.Space, mode string) []cell {
	if mode == fun.Minimum {
		sort.SliceStable(n.antibodies, func(i, j int) bool {
			return n.antibodies[i].affinity < n.antibodies[j].affinity
		})
	} else {
		sort.SliceStable(n.antibodies, func(i, j int) bool {
			return n.antibodies[i].affinity > n.antibodies[j].affinity
		})
	}

	clones := make([]cell, 0)
	for i := 0; i < n.bestQuantity; i++ {
		for j := 0; j < n.cloneQuantity; j++ {
			cell := n.antibodies[i].clone()
			cell.mutate(n.mutationRate, s, f)
			clones = append(clones, cell)
		}
	}

	if mode == fun.Minimum {
		sort.SliceStable(clones, func(i, j int) bool {
			return clones[i].affinity < clones[j].affinity
		})
	} else {
		sort.SliceStable(clones, func(i, j int) bool {
			return clones[i].affinity > clones[j].affinity
		})
	}

	return clones
}

func (n *network) unite(clones []cell, f fun.TargetFunc, s mat.Space, mode string) {
	for _, c := range clones {
		n.antibodies = append(n.antibodies, c)
	}

	if mode == fun.Minimum {
		sort.SliceStable(n.antibodies, func(i, j int) bool {
			return n.antibodies[i].affinity < n.antibodies[j].affinity
		})
	} else {
		sort.SliceStable(n.antibodies, func(i, j int) bool {
			return n.antibodies[i].affinity > n.antibodies[j].affinity
		})
	}

	n.antibodies = n.antibodies[:n.population-2]
	n.antibodies = append(n.antibodies, generateCell(f, s))
	n.antibodies = append(n.antibodies, generateCell(f, s))
}
