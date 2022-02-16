package section8

import (
	"log"
	"time"
)

func DemoTrace() {
	defer bigSlowOperation("bigslowoperation")()
	time.Sleep(5 * time.Second) //模拟慢操作
}

func bigSlowOperation(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() { log.Printf("exit %s: (%s)\n", msg, time.Since(start)) }
}
