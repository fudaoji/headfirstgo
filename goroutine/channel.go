package goroutine

import (
	"log"
	"time"
)

func reportNap(name string, times int) {
	for i := 0; i < times; i++ {
		log.Println(name, " sleeping")
		time.Sleep(1 * time.Second)
	}
	log.Println(name, " wakes up!")
}

func send(myChan chan string) {
	reportNap("send goroutine", 2)
	log.Println("======sending value=======")
	myChan <- "a"
	log.Println("======sending value=======")
	log.Println("======sending value=======")
	myChan <- "b"
}

func TestChannel() {
	myChan := make(chan string, 1)
	go send(myChan)
	reportNap("receive goroutine", 5)
	log.Println(<-myChan)
	log.Println("======sending value=======")
	log.Println("======sending value=======")
	log.Println(<-myChan)
}
