package section4

import (
	"fmt"
)

func DemoScaleBy() {
	a := Point{10, 30}
	scaleby := (*Point).ScaleBy
	scaleby(&a, 2)
	fmt.Printf("scaleby: %T\n", scaleby)
	fmt.Println(a)
}

//DemoDistance 方法表达式
func DemoDistance() {
	a := Point{10, 30}
	b := Point{25, 45}
	distance := Point.Distance
	fmt.Printf("the distance between a and b is %.4f\n", distance(a, b))
}
