package section2

import "fmt"

func DemoAppend3() {
	var y []int
	y = appendIntI(y, 1)
	y = appendIntI(y, 2, 3, 4)
	fmt.Println(y)
}

func DemoAppend2() {
	var y []int
	x := []int{1, 2, 3}
	y = append(y, x...)
	fmt.Println(y)
}

func DemoAppend1() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(y), y)
		x = y
	}
}

func DemoAppend() {
	s := []int{1, 2}
	s = appendInt(s, 3)
	s = appendInt(s, 4)
	fmt.Println(s)
}

func appendIntI(s []int, n ...int) []int {
	var z []int

	sLen := len(s)
	zLen := sLen + len(n)

	if zLen <= cap(s) {
		z = s[:zLen] //空间足够，扩展一个元素
	} else {
		//容量不够，扩展容量再推入
		zCap := zLen
		if zCap < 2*sLen {
			zCap = 2 * sLen
		}
		z = make([]int, zLen, zCap)
		copy(z, s)
	}
	copy(z[sLen:], n) //多个元素，所以用copy
	return z
}

func appendInt(s []int, n int) []int {
	var z []int
	//如果容量足够，直接推入
	//容量不够，扩展容量再推入
	sLen := len(s)
	zLen := sLen + 1

	if zLen <= cap(s) {
		z = s[:zLen] //空间足够，扩展一个元素
	} else {
		zCap := zLen
		if zCap < 2*sLen {
			zCap = 2 * sLen
		}
		z = make([]int, zLen, zCap)
		copy(z, s)
	}
	z[sLen] = n
	return z
}
