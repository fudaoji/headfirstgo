package section1

import (
	"crypto/sha256"
	"fmt"
)

func Demo() {
	arr := [...]string{0: "h", 1: "e"}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	arr2 := [...]int{99: -1}
	fmt.Println(arr2)
}

func Demosha256() {
	fmt.Printf("%T\n", sha256.Sum256([]byte("x")))
}

func SliceDemo() {
	s := make([]int, 3, 100)
	s1 := []int(nil)
	fmt.Println(s)
	fmt.Println(s1 == nil)
}

func SliceDemo2() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[2:5]
	s1 := a[6:9]
	s = append(s, 10, 9)
	a[3] = 0
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s1)

	s2 := make([]int, 3, 4)
	s2[1] = 2
	s3 := s2[:2]
	s3 = append(s3, 10)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
}
