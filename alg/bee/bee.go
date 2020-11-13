package bee

import (
	"fmt"

	"github.com/Elfsilon/opt/utils/fun"
	"github.com/Elfsilon/opt/utils/mat"
)

func newBee() bee {
	return bee{}
}

type bee struct {
	position mat.Vec2
	localExt mat.Extremum
}

func (b *bee) setPosition(v mat.Vec2) {
	b.position = v
}

func (b *bee) evalValue(f fun.TargetFunc) {
	b.localExt = mat.NewExtremum(b.position, f.Eval(b.position))
}

func (b *bee) work(position mat.Vec2, f fun.TargetFunc) {
	b.setPosition(position)
	b.evalValue(f)
}

func (b *bee) String() string {
	return fmt.Sprintf("Bee: position = %-.2f extremum coords = %-.4f extremum value = %-.2v\n", b.position, b.localExt.Coord, b.localExt.Value)
}
