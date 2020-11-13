package swarm

import (
	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

// particle ...
type particle struct {
	position mat.Vec2
	velocity mat.Vec2
	localExt mat.Extremum
}

// newParticle ...
func newParticle(position, velocity mat.Vec2, ext mat.Extremum) particle {
	return particle{
		position: position,
		velocity: velocity,
		localExt: ext,
	}
}

func (p *particle) evalExtremum(mode string, f fun.TargetFunc) {
	switch mode {
	case "maximum":
		p.evalMaximum(f)
	case "minimum":
		p.evalMinimum(f)
	}
}

func (p *particle) evalMaximum(f fun.TargetFunc) {
	value := f.Eval(p.position)
	if value > p.localExt.Value {
		ext := mat.NewExtremum(p.position, value)
		p.localExt = ext
	}
}

func (p *particle) evalMinimum(f fun.TargetFunc) {
	value := f.Eval(p.position)
	if value < p.localExt.Value {
		ext := mat.NewExtremum(p.position, value)
		p.localExt = ext
	}
}

func (p *particle) correctVelocity(phi1, phi2 float64, globalExt mat.Extremum) {
	r1 := utils.RandFloat(0.0001, 0.9999)
	r2 := utils.RandFloat(0.0001, 0.9999)

	part1 := p.localExt.Coord.Sub(p.position)
	part1 = part1.Mult(r1 * phi1)

	part2 := globalExt.Coord.Sub(p.position)
	part2 = part2.Mult(r2 * phi2)

	sum := part1.Add(part2)

	p.velocity = p.velocity.Add(sum)
}

func (p *particle) correctPosition() {
	p.position = p.position.Add(p.velocity)
}
