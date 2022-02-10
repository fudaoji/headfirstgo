package section3

func incr(p *int) int {
	*p++
	return *p
}

type fl1 float64
type fl2 float64

func Demo2() {
	var i1 fl1 = 1.0
	var i2 fl2 = 2.0

	println(i1 - 1)
	println(i2 - 1)
}

func init() {
	println("ch2/section3 init")
}

func Demo1() {
	num := 1
	incr(&num)
	println(incr(&num))
}
