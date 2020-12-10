package bact

import (
	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newBacterium() bacterium {
	return bacterium{}
}

type bacterium struct {
	position    mat.Vec2
	direction   mat.Vec2
	stepsize    float64
	velocity    mat.Vec2
	value       float64
	healthState float64
}

func (b *bacterium) setStep(size float64) {
	b.stepsize = size
}

func (b *bacterium) setDirection(dir mat.Vec2) {
	b.direction = dir
}

func (b *bacterium) setPosition(pos mat.Vec2) {
	b.position = pos
}

func (b *bacterium) randPosition(f fun.TargetFunc, s mat.Space) {
	x := utils.RandFloat(s.Xmin, s.Xmax)
	y := utils.RandFloat(s.Ymin, s.Ymax)

	b.position = mat.NewVec2(x, y)
	b.value = f.Eval(b.position)

	b.flipDirection()
}

func (b *bacterium) initOnExtremum(stepsize float64, f fun.TargetFunc, ext mat.Extremum, distributionRate float64) {
	x := utils.RandFloat(ext.Coord.X-distributionRate, ext.Coord.X+distributionRate)
	y := utils.RandFloat(ext.Coord.Y-distributionRate, ext.Coord.Y+distributionRate)

	b.position = mat.NewVec2(x, y)
	b.value = f.Eval(b.position)
	b.healthState += b.value

	b.stepsize = stepsize

	dx := utils.RandFloat(-1, 1)
	dy := utils.RandFloat(-1, 1)
	b.direction = mat.NewVec2(dx, dy)

	dirNorm := mat.ENorm(b.direction)

	b.velocity = b.direction.Mult(b.stepsize / dirNorm)
}

func (b *bacterium) init(stepsize float64, f fun.TargetFunc, s mat.Space) {
	x := utils.RandFloat(s.Xmin, s.Xmax)
	y := utils.RandFloat(s.Ymin, s.Ymax)

	b.position = mat.NewVec2(x, y)
	b.value = f.Eval(b.position)
	b.healthState += b.value

	b.stepsize = stepsize

	dx := utils.RandFloat(-1, 1)
	dy := utils.RandFloat(-1, 1)
	b.direction = mat.NewVec2(dx, dy)

	dirNorm := mat.ENorm(b.direction)

	b.velocity = b.direction.Mult(b.stepsize / dirNorm)
}

func (b *bacterium) flipDirection() {
	dx := utils.RandFloat(-1, 1)
	dy := utils.RandFloat(-1, 1)

	b.direction = mat.NewVec2(dx, dy)
	dirNorm := mat.ENorm(b.direction)

	b.velocity = b.direction.Mult(b.stepsize / dirNorm)
}

func (b *bacterium) move(f fun.TargetFunc, mode string) {
	newPosition := b.position.Add(b.velocity)
	newValue := f.Eval(newPosition)

	if mode == fun.Maximum {
		if newValue > b.value {
			b.position = newPosition
			b.value = newValue
			b.healthState += newValue
		} else {
			b.flipDirection()
		}
	} else {
		if newValue < b.value {
			b.position = newPosition
			b.value = newValue
			b.healthState += newValue
		} else {
			b.flipDirection()
		}
	}

}

func (b *bacterium) hemotaxis(lim int, f fun.TargetFunc, mode string) {
	oldValue := b.value
	for i := 0; i < lim; i++ {
		b.move(f, mode)
		if b.value < oldValue {
			break
		}
		oldValue = b.value
	}
}
