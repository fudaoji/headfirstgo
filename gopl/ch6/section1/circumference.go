package section1

type Path []Point

//Distance Get the distance between p and q
func (p Path) Circumference() float64 {
	sum := 0.0
	for i, dot := range p {
		if i > 0 {
			sum += dot.Distance(p[i-1])
		}
	}
	return sum
}
