package bee

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// ArtificialBeeColonyAlg ...
type ArtificialBeeColonyAlg struct {
	TFunc          *fun.TargetFunc
	Hive           *hive
	Space          *mat.Space
	Mode           string
	IterationsLim  int
	Iteration      int
	Scouts         int
	Workers        int
	EliteAreas     int
	PerspectAreas  int
	EliteBee       int
	PerspectBee    int
	EliteRadius    float64
	PerspectRadius float64
}

// NewArtificialBeeColonyAlg ...
func NewArtificialBeeColonyAlg(opt *Options) (*ArtificialBeeColonyAlg, error) {
	err := opt.validate()
	if err != nil {
		return nil, err
	}

	workersCount := opt.EliteBee*opt.Elite + opt.PerspectBee*opt.Perspect
	scoutCount := opt.Elite + opt.Perspect
	hive := newHive(scoutCount, workersCount, opt.Elite, opt.Perspect, opt.EliteBee, opt.PerspectBee, opt.EliteRadius, opt.PerspectRadius)

	return &ArtificialBeeColonyAlg{
		TFunc:         &opt.TargetFunc,
		Hive:          &hive,
		Space:         &opt.Space,
		Mode:          opt.Mode,
		IterationsLim: opt.IterationsLim,
		Iteration:     0,
		Workers:       workersCount,
		Scouts:        scoutCount,
		EliteAreas:    opt.Elite,
		PerspectAreas: opt.Perspect,
		EliteBee:      opt.EliteBee,
		PerspectBee:   opt.PerspectBee,
	}, nil
}

// Start ...
func (a *ArtificialBeeColonyAlg) Start() (mat.Extremum, string) {
	rand.Seed(time.Now().UnixNano())

	for {
		a.Hive.explore(*a.TFunc, *a.Space)
		a.Hive.sendWorkers(*a.TFunc)
		a.Hive.updateAreas()
		a.Hive.sortAreas(a.Mode)
		a.Hive.updateExtremums()

		a.Iteration++

		if a.Iteration >= a.IterationsLim {
			break
		}
	}

	return a.Hive.globalExt[0], a.String()
}

func (a *ArtificialBeeColonyAlg) String() string {
	return fmt.Sprint("ALGORITHM STATS\n") +
		fmt.Sprintf("Iterations limit: %v\n", a.IterationsLim) +
		fmt.Sprintf("Current iteration: %v\n", a.Iteration) +
		fmt.Sprintf("Searching space: %v\n", *a.Space) +
		fmt.Sprintf("Mode: %v\n", a.Mode) +
		fmt.Sprintf("Elite areas number: %v\n", a.EliteAreas) +
		fmt.Sprintf("Prespective areas number: %v\n", a.PerspectAreas) +
		fmt.Sprintf("Elite bees number: %v\n", a.EliteBee) +
		fmt.Sprintf("Prespective bees number: %v\n", a.PerspectBee) +
		fmt.Sprintf("Worker bees number: %v\n", a.Workers) +
		fmt.Sprintf("Scout bees number: %v\n", a.Scouts) +
		fmt.Sprintf("HIVE: \n%v\n", a.Hive.String())
}
