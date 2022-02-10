package goroutine

import (
	"log"
	"time"
)

func a(channel chan int) {
	var i int
	for i = 0; i <= 100; i++ {
		log.Print("a")
	}
	channel <- i
}

func b() {
	for i := 0; i <= 100; i++ {
		log.Print("b")
	}
}

func c() {
	for i := 0; i <= 100; i++ {
		log.Print("c")
	}
}

func Test() {
	aRes := make(chan int)
	go a(aRes)
	go b()
	go c()
	time.Sleep(5 * time.Second)
	println(<-aRes)
}
