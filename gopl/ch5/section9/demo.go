package section9

import (
	"fmt"
	"os"
	"runtime"
)

func DemoPrintStack() {
	defer printStack()
	f(4)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func Demof() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer f(%d)\n", x)
	f(x - 1)
}
