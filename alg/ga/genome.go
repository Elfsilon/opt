package ga

import (
	"github.com/Elfsilon/opt/utils/fun"
	"math/rand"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/mat"
)

func newGenome(position mat.Vec2, value float64) genome {
	return genome{
		position: position,
		value:    value,
	}
}

// Genome ...
type genome struct {
	position mat.Vec2
	value    float64
}

func (g *genome) mutate(probability, rate float64, f fun.TargetFunc) {
	rx, ry := utils.RandFloat(0, 1), utils.RandFloat(0, 1)
	dirx, diry := rand.Intn(2), rand.Intn(2)

	if rx < probability {
		if dirx == 0 {
			g.position.X += rate
		} else {
			g.position.X -= rate
		}
	}

	if ry < probability {
		if diry == 0 {
			g.position.Y += rate
		} else {
			g.position.Y -= rate
		}
	}

	g.value = f.Eval(g.position)
}
