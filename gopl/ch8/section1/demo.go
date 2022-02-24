package section1

import (
	"fmt"
	"time"
)

//DemoGoroutine 使用goroutine演示fib
func DemoGoroutine() {
	go spinner(100 * time.Millisecond)
	n := 45
	fibN := fib(n)
	fmt.Printf("the fib of %d is: %d", n, fibN)
}

//fib
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

//spinner 进度加载条
func spinner(t time.Duration) {
	for { //生命周期有主goroutine决定
		for _, r := range `_\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(t)
		}
	}
}
