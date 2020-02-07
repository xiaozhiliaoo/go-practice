package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

//函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy2(factor float64) {
	p.X *= factor
	p.Y *= factor
}
