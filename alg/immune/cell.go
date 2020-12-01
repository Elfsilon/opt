package immune

import (
	// "math"

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
	dx := utils.RandFloat(-0.5, 0.5) * intensity
	dy := utils.RandFloat(-0.5, 0.5) * intensity

	c.coords.Add(mat.NewVec2(dx, dy))
	c.affinity = f.Eval(c.coords)
}
