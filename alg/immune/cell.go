package immune

import (
	// "math/rand"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func generateCell(f fun.TargetFunc, s mat.Space) cell {
	x := utils.RandFloat(s.Xmin, s.Xmax)
	y := utils.RandFloat(s.Ymin, s.Ymax)
	coords := mat.NewVec2(x, y)
	return cell{
		coords:   coords,
		affinity: f.Eval(coords),
	}
}

func generateCellOnExtremum(f fun.TargetFunc, ext mat.Extremum, distributionRate float64) cell {
	x := utils.RandFloat(ext.Coord.X-distributionRate, ext.Coord.X+distributionRate)
	y := utils.RandFloat(ext.Coord.Y-distributionRate, ext.Coord.Y+distributionRate)
	coords := mat.NewVec2(x, y)
	return cell{
		coords:   coords,
		affinity: f.Eval(coords),
	}
}

func newCell(coords mat.Vec2, affinity float64) cell {
	return cell{
		coords:   coords,
		affinity: affinity,
	}
}

type cell struct {
	coords   mat.Vec2
	affinity float64
}

func (c *cell) clone() cell {
	return newCell(c.coords, c.affinity)
}

func (c *cell) mutate(intensity float64, s mat.Space, f fun.TargetFunc) {
	dx := float64(utils.RandInt(-1, 1)) * intensity
	dy := float64(utils.RandInt(-1, 1)) * intensity

	c.coords.Add(mat.NewVec2(dx, dy))
	c.affinity = f.Eval(c.coords)
}
