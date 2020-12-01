package ga

import (
	"math"
	"math/rand"
	"sort"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newEvolution(populationSize int, mutationRate, mutationProbability float64) evolution {
	return evolution{
		population:          make([]genome, populationSize),
		mutationRate:        mutationRate,
		mutationProbability: mutationProbability,
	}
}

type evolution struct {
	population          []genome
	mutationRate        float64
	mutationProbability float64
}

func (e *evolution) init(f fun.TargetFunc, s mat.Space) {
	for i := range e.population {
		x := utils.RandFloat(s.Xmin, s.Xmax)
		y := utils.RandFloat(s.Ymin, s.Ymax)

		e.population[i].position = mat.NewVec2(x, y)
		e.population[i].value = f.Eval(e.population[i].position)
	}
}

func (e *evolution) selection(f fun.TargetFunc, mode string) []genome {
	if mode == fun.Minimum {
		sort.SliceStable(e.population, func(i, j int) bool {
			return e.population[i].value < e.population[j].value
		})
	} else {
		sort.SliceStable(e.population, func(i, j int) bool {
			return e.population[i].value > e.population[j].value
		})
	}

	return e.population[:len(e.population)/2]
}

func (e *evolution) crossover(selected []genome, f fun.TargetFunc, mode string) {
	var nextGeneration []genome

	for len(nextGeneration) < len(e.population) {
		// Select 2 parents
		firstParent := rand.Intn(len(selected))
		var secondParent int
		for {
			secondParent = rand.Intn(len(selected))
			if firstParent != secondParent {
				break
			}
		}

		x1 := selected[firstParent].position.X
		y1 := selected[firstParent].position.Y
		x2 := selected[secondParent].position.X
		y2 := selected[secondParent].position.Y

		// Linear crossover
		px := []float64{
			(x1 + x2) / 2,
			(3*x1 - x2) / 2,
			-(x1 - 3*x2) / 2,
		}
		py := []float64{
			(y1 + y2) / 2,
			(3*y1 - y2) / 2,
			-(y1 - 3*y2) / 2,
		}

		var fbest float64
		var xbest, ybest float64 = -10000, -10000
		if mode == fun.Minimum {
			fbest = math.MaxFloat32
		} else {
			fbest = -math.MaxFloat32
		}

		for _, pxEl := range px {
			for _, pyEl := range py {
				pos := mat.NewVec2(pxEl, pyEl)
				val := f.Eval(pos)

				if mode == fun.Minimum {
					if val < fbest {
						fbest = val
						xbest, ybest = pxEl, pyEl
					}
				} else {
					if val > fbest {
						fbest = val
						xbest, ybest = pxEl, pyEl
					}
				}
			}
		}

		gen := newGenome(mat.NewVec2(xbest, ybest), fbest)
		nextGeneration = append(nextGeneration, gen)
	}

	e.population = nextGeneration
}

func (e *evolution) mutation(f fun.TargetFunc, mode string) {
	for i := range e.population {
		e.population[i].mutate(e.mutationProbability, e.mutationRate, f)
	}

	if mode == fun.Minimum {
		sort.SliceStable(e.population, func(i, j int) bool {
			return e.population[i].value < e.population[j].value
		})
	} else {
		sort.SliceStable(e.population, func(i, j int) bool {
			return e.population[i].value > e.population[j].value
		})
	}
}
