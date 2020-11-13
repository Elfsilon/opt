package mat

import (
	"math"
)

// NewVec2 - Vec2 constructor function
func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

// Vec2 - representation of 2D vector
type Vec2 struct {
	X, Y float64
}

// Magnitude - length of vector
func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Mult - vector multiplication by scalar
func (v *Vec2) Mult(k float64) Vec2 {
	return NewVec2(v.X*k, v.Y*k)
}

// Add - vector addition
func (v *Vec2) Add(v2 Vec2) Vec2 {
	return NewVec2(v.X+v2.X, v.Y+v2.Y)
}

// Sub - vector substraction
func (v *Vec2) Sub(v2 Vec2) Vec2 {
	return NewVec2(v.X-v2.X, v.Y-v2.Y)
}

// ENorm - Euclidean norm of vector 2d
func ENorm(v Vec2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
