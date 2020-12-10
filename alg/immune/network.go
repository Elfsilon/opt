package immune

import (
	"math"
	"math/rand"
	"sort"

	"github.com/Elfsilon/opt/utils"
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

func (n *network) generateAntibodiesOnExtremum(f fun.TargetFunc, ext mat.Extremum, distrRatio float64) {
	for i := 0; i < n.population; i++ {
		cell := generateCellOnExtremum(f, ext, distrRatio)
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

	n.antibodies = n.antibodies[:n.population]
}

func (n *network) scatter(probability float64, count int, f fun.TargetFunc, s mat.Space) {
	for i := 0; i < count; i++ {
		r := utils.RandFloat(0, 1)
		if r < probability {
			index := rand.Intn(len(n.antibodies))
			n.antibodies[index] = generateCell(f, s)
		}
	}
}

func (n *network) selectBestSolution(mode string) mat.Extremum {
	var bestSol mat.Extremum

	if mode == fun.Maximum {
		bestSol.Value = -math.MaxFloat32
	} else {
		bestSol.Value = math.MaxFloat32
	}

	for _, ab := range n.antibodies {
		if mode == fun.Maximum {
			if ab.affinity > bestSol.Value {
				bestSol = mat.NewExtremum(ab.coords, ab.affinity)
			}
		} else {
			if ab.affinity < bestSol.Value {
				bestSol = mat.NewExtremum(ab.coords, ab.affinity)
			}
		}
	}

	return bestSol
}
