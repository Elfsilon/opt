package swarm

import (
	"math"

	"github.com/Elfsilon/opt/utils"
	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

type swarm struct {
	swarm     []particle
	size      int
	globalExt mat.Extremum
}

func generateSwarm(swarmsize int, space mat.Space, mode string) swarm {
	swarm := swarm{}
	swarm.generate(swarmsize, space, mode)
	return swarm
}

func (s *swarm) generate(swarmsize int, space mat.Space, mode string) {
	s.size = swarmsize
	for i := 0; i < swarmsize; i++ {
		x := utils.RandFloat(space.Xmin, space.Xmax)
		y := utils.RandFloat(space.Ymin, space.Ymax)

		startPos := mat.NewVec2(x, y)
		startV := mat.NewVec2(utils.RandFloat(0, 1), utils.RandFloat(0, 1))

		var startExt mat.Extremum
		switch mode {
		case fun.Maximum:
			startExt = mat.NewExtremum(mat.NewVec2(0, 0), -math.MaxFloat32)
		case fun.Minimum:
			startExt = mat.NewExtremum(mat.NewVec2(0, 0), math.MaxFloat32)
		}

		part := newParticle(startPos, startV, startExt)
		s.swarm = append(s.swarm, part)
	}
	s.globalExt = s.swarm[0].localExt
}

// Calculates all local extremums of partcles then
// find best of them and memorize it in swarm.globalExt
func (s *swarm) evalExtremums(mode string, f fun.TargetFunc) {
	for i := range s.swarm {
		s.swarm[i].evalExtremum(mode, f)
		switch mode {
		case fun.Maximum:
			if s.swarm[i].localExt.Value > s.globalExt.Value {
				s.globalExt = s.swarm[i].localExt
			}
		case fun.Minimum:
			if s.swarm[i].localExt.Value < s.globalExt.Value {
				s.globalExt = s.swarm[i].localExt
			}
		}
	}
}

// Calculates new velocity value of all swarm's particles
// and moves them to new coords = current coords + new velocity
func (s *swarm) correctParticles(phi1, phi2 float64) {
	for i := range s.swarm {
		s.swarm[i].correctVelocity(phi1, phi2, s.globalExt)
		s.swarm[i].correctPosition()
	}
}
