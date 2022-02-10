package section2

import (
	"fmt"
	"os"
	"time"
)

//demo
func Demo() {
	var s, sep string
	for _, item := range os.Args[1:] {
		s += sep + item
		sep = "\n"
	}
	println(s)
}

//exer1 excecise 1
func Exer1() {
	println(os.Args[0])
}

//exer2 excecise 2
func Exer2() {
	var s string
	for i, item := range os.Args[1:] {
		s += fmt.Sprintf("%d:%s\n", i, item)
	}
	println(s)
}

//exer3 excecise 3
func Exer3() {
	println(time.Now().Local().UnixMicro())
	Demo()
	println(time.Now().Local().UnixMicro())
}
