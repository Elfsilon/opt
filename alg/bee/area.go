package bee

import (
	"fmt"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/mat"
)

type area struct {
	center mat.Vec2
	radius float64
	value  float64
}

func newArea(center mat.Vec2, value, radius float64) area {
	return area{
		center: center,
		value:  value,
		radius: radius,
	}
}

func (a *area) setRadius(radius float64) {
	a.radius = radius
}

func (a *area) randomPoint() mat.Vec2 {
	xmin, xmax := a.center.X-a.radius, a.center.X+a.radius
	ymin, ymax := a.center.Y-a.radius, a.center.Y+a.radius

	x := utils.RandFloat(xmin, xmax)
	y := utils.RandFloat(ymin, ymax)

	return mat.NewVec2(x, y)
}

func (a *area) String() string {
	return fmt.Sprintf("Area: center = %-.2f value = %-.4f radius = %-.2v \n", a.center, a.value, a.radius)
}
