package gen

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/alg/utils"
)

// Space ...
type Space struct {
	Xmin, Xmax, Ymin, Ymax float64
}

// NewSpace ...
func NewSpace(xmin, xmax, ymin, ymax float64) Space {
	return Space{
		Xmin: xmin,
		Xmax: xmax,
		Ymin: ymin,
		Ymax: ymax,
	}
}

// FitnessFunc ...
type FitnessFunc interface {
	Eval(g Genome) float64
	EvalGen(p Generation) ([]float64, []float64)
}

// RozenbrocFunc ...
type RozenbrocFunc struct{}

// Eval ...
func (r *RozenbrocFunc) Eval(g Genome) float64 {
	return (1-g.x)*(1-g.x) + 100*(g.y-g.x*g.x)*(g.y-g.x*g.x)
}

// EvalGen ...
func (r *RozenbrocFunc) EvalGen(g Generation) ([]float64, []float64) {
	var (
		values []float64
		sum    float64
	)
	for _, genome := range g {
		val := r.Eval(genome)
		values = append(values, val)
		sum += val
	}

	var percents []float64
	for _, val := range values {
		percents = append(percents, val/sum)
	}

	return values, percents
}

// Genome ...
type Genome struct {
	x, y float64
}

func (g *Genome) mutate(probability, step float64, s Space) {
	if rand.Float64() < probability {
		dir := rand.Intn(2)
		if dir == 0 {
			g.x += step
		} else {
			g.x -= step
		}
	}
	if rand.Float64() < probability {
		dir := rand.Intn(2)
		if dir == 0 {
			g.y += step
		} else {
			g.y -= step
		}
	}
}

// NewGenome ...
func NewGenome(x, y float64) Genome {
	return Genome{x: x, y: y}
}

// Generation ...
type Generation []Genome

func (p *Generation) generate(n int, s Space) {
	for i := 0; i < n; i++ {
		x := utils.RandFloat(s.Xmin, s.Xmax)
		y := utils.RandFloat(s.Ymin, s.Ymax)
		*p = append(*p, NewGenome(x, y))
	}
}

func (p *Generation) mindist(rnd float64, percents []float64) (float64, int) {
	var mindist float64 = math.MaxFloat32
	var index int = 0
	for i, p := range percents {
		dist := math.Abs(rnd - p)
		if dist < mindist {
			index = i
			mindist = dist
		}
	}
	return mindist, index
}

func (p *Generation) selection(f FitnessFunc, weights []float64) Pairs {
	var pairs Pairs

	var meanw float64 = 0
	for _, w := range weights {
		meanw += w
	}
	meanw /= float64(len(weights))

	var selected Generation
	for _, g := range *p {
		if f.Eval(g) < meanw {
			selected = append(selected, g)
		}
	}

	for len(pairs) != len(*p) {
		var pair [2]Genome
		i1 := rand.Intn(len(selected))
		for {
			i2 := rand.Intn(len(selected))
			if i1 != i2 {
				pair = [2]Genome{(*p)[i1], (*p)[i2]}
				break
			}
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (p *Generation) mutation(probability, step float64, s Space) {
	for i := range *p {
		(*p)[i].mutate(probability, step, s)
	}
}

// Pairs ...
type Pairs [][2]Genome

func (p *Pairs) crossover(f FitnessFunc) Generation {
	var newGeneration Generation
	for i := 0; i < len(*p); i++ {
		x1, y1 := (*p)[i][0].x, (*p)[i][0].y
		x2, y2 := (*p)[i][1].x, (*p)[i][1].y

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

		var fbest, xbest, ybest float64 = math.MaxFloat32, -1, -1
		for _, i := range px {
			for _, j := range py {
				if z := f.Eval(NewGenome(i, j)); z < fbest {
					fbest = z
					xbest, ybest = i, j
				}
			}
		}

		newGeneration = append(newGeneration, NewGenome(float64(xbest), float64(ybest)))
	}
	return newGeneration
}

// Evolution ...
func Evolution() {
	fitnessFunc := RozenbrocFunc{}
	space := NewSpace(-10, 10, -10, 10)
	var e float64 = 0.1
	generationsCount := 10000

	var g Generation
	g.generate(5, space)

	flag := true
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < generationsCount && flag; i++ {
		weights, _ := fitnessFunc.EvalGen(g)
		pairs := g.selection(&fitnessFunc, weights)
		g := pairs.crossover(&fitnessFunc)
		g.mutation(0.05, 0.15, space)
		for _, genome := range g {
			if fitnessFunc.Eval(genome) <= e {
				printGen(g, "New generation after mutation", &fitnessFunc)
				fmt.Println("Ended", genome, i)
				flag = false
				break
			}
		}
	}
	// printGen(g, "New generation after mutation", &fitnessFunc)
}

func printGen(g Generation, text string, f FitnessFunc) {
	fmt.Println(text)
	for _, el := range g {
		fmt.Println(el, "Z =", f.Eval(el))
	}
	fmt.Println("")
}
