package section4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fudaoji/learn-go/gopl/ch8/section2"
)

func DemoChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	if len(ch) < cap(ch) {
		ch <- 3
	} else {
		fmt.Println("通道已满")
	}
}

func DemoPipeline2() {
	//三个管道  natures、squares、printer
	natures := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			natures <- i
			time.Sleep(1 * time.Second)
		}
		close(natures)
	}()

	go func() {
		for x := range natures {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

func DemoPipeline1() {
	//三个管道  natures、squares、printer
	natures := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i >= 0; i++ {
			natures <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			x := <-natures
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}

func DemoNetcat3() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		ch <- 1
	}()
	section2.MustCopy(conn, os.Stdin)
	<-ch
}

//DemoClock
func DemoReverb1() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如链接中止
			continue
		}
		go handleConn(conn)
	}
}

//handleConn 连接处理
func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
