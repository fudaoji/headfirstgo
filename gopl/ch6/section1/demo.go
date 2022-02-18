package section1

import (
	"fmt"
	"log"
)

func DemoScaleBy() {
	a := Point{10, 30}
	log.Printf("%p", &a)
	a.ScaleBy(2)

	fmt.Println(a)
}

func DemoCircumference() {
	p1 := Point{1, 2}
	p2 := Point{5, 2}
	p3 := Point{5, 4}

	path2 := Path{
		p1,
		p2,
		p3,
		p1,
	}

	path := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("the circumference of path is %.4f\n", path.Circumference())
	fmt.Printf("the circumference of path2 is %.4f\n", path2.Circumference())
}

func DemoDistance() {
	a := Point{10, 30}
	b := Point{25, 45}
	fmt.Printf("the distance between a and b is %.4f\n", a.Distance(b))
}
