package mat

// Extremum - extremum point struct
type Extremum struct {
	Coord Vec2
	Value float64
}

// NewExtremum ...
func NewExtremum(coord Vec2, value float64) Extremum {
	return Extremum{Coord: coord, Value: value}
}

// Space ...
type Space struct {
	Xmin, Xmax, Ymin, Ymax float64
}

// NewSpace ...
func NewSpace(Xmin, Xmax, Ymin, Ymax float64) Space {
	return Space{Xmin: Xmin, Xmax: Xmax, Ymin: Ymin, Ymax: Ymax}
}
