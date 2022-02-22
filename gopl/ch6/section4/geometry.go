package section4

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

//Distance Get the distance between p and q
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor int) {
	p.X *= float64(factor)
	p.Y *= float64(factor)
}
