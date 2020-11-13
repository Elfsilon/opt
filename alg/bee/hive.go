package bee

import (
	"fmt"
	"sort"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newHive(
	scoutsCount,
	workersCount,
	eliteCount,
	perspectCount,
	eliteBeesCount,
	persectBeesCount int,
	eliteRadius,
	prespectRadius float64,
) hive {
	scouts := make([]bee, scoutsCount)
	workers := make([]bee, workersCount)
	return hive{
		scouts:           scouts,
		workers:          workers,
		eliteCount:       eliteCount,
		perspectCount:    perspectCount,
		eliteBeesCount:   eliteBeesCount,
		persectBeesCount: persectBeesCount,
		eliteRadius:      eliteRadius,
		perspectRadius:   prespectRadius,
	}
}

type hive struct {
	globalExt []mat.Extremum
	scouts    []bee
	workers   []bee
	areas     []area

	eliteCount       int
	perspectCount    int
	eliteBeesCount   int
	persectBeesCount int
	eliteRadius      float64
	perspectRadius   float64
}

// Sends worker bees to avaliable areas
// First cycle - sending workers to elite areas
// Second - sending workers to perspective areas
func (h *hive) sendWorkers(f fun.TargetFunc) {
	for i := 0; i < h.eliteCount; i++ {
		h.areas[i].setRadius(h.eliteRadius)
		startJ := i * h.eliteBeesCount
		for j := startJ; j < startJ+h.eliteBeesCount; j++ {
			point := h.areas[i].randomPoint()
			h.workers[j].work(point, f)
		}
	}

	for i := h.eliteCount; i < h.eliteCount+h.perspectCount; i++ {
		h.areas[i].setRadius(h.perspectRadius)
		startJ := h.eliteBeesCount*h.eliteCount + (i-h.eliteCount)*h.persectBeesCount
		for j := startJ; j < startJ+h.persectBeesCount; j++ {
			point := h.areas[i].randomPoint()
			h.workers[j].work(point, f)
		}
	}
}

// Sends scout bees which explores new avaliable areas
func (h *hive) explore(f fun.TargetFunc, s mat.Space) {
	for i := range h.scouts {
		x := utils.RandFloat(s.Xmin, s.Xmax)
		y := utils.RandFloat(s.Ymin, s.Ymax)

		h.scouts[i].work(mat.NewVec2(x, y), f)

		area := newArea(h.scouts[i].position, h.scouts[i].localExt.Value, -1)
		h.areas = append(h.areas, area)
	}
}

// Reconstructs []areas with new data
func (h *hive) updateAreas() {
	var areas []area
	for _, s := range h.scouts {
		areas = append(areas, newArea(s.position, s.localExt.Value, -1))
	}
	for _, w := range h.workers {
		areas = append(areas, newArea(w.position, w.localExt.Value, -1))
	}
	h.areas = areas
}

// Sorts areas by desc/asc mode
func (h *hive) sortAreas(mode string) {
	if mode == fun.Maximum {
		sort.SliceStable(h.areas, func(i, j int) bool {
			return h.areas[i].value > h.areas[j].value
		})
	}
	if mode == fun.Minimum {
		sort.SliceStable(h.areas, func(i, j int) bool {
			return h.areas[i].value < h.areas[j].value
		})
	}
}

// Reconstructs globalExt with data about updated areas
func (h *hive) updateExtremums() {
	var globalExt []mat.Extremum
	for i := 0; i < h.eliteCount; i++ {
		globalExt = append(globalExt, mat.NewExtremum(h.areas[i].center, h.areas[i].value))
	}
	h.globalExt = globalExt
}

func (h *hive) String() string {
	extremums := "Extremums:\n"
	for _, ext := range h.globalExt {
		extremums += fmt.Sprintf("  Extremum coords = %-.2v value = %-.4v\n", ext.Coord, ext.Value)
	}

	scouts := "Scouts:\n"
	for _, s := range h.scouts {
		scouts += fmt.Sprintf("  %v", s.String())
	}

	workers := "Workers:\n"
	for _, w := range h.workers {
		workers += fmt.Sprintf("  %v", w.String())
	}

	areas := "Areas:\n"
	for _, a := range h.areas {
		areas += fmt.Sprintf("  %v", a.String())
	}

	return extremums + areas + scouts + workers
}
