package section2

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

//DemoClockGoroutine 使用并发演示clock
func DemoClockGoroutine() {
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

func DemoNetcat() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	MustCopy(os.Stdout, conn)
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//DemoClock
func DemoClock() {
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
		handleConn(conn)
	}
}

//handleConn 连接处理
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006/1/2 15:04:05\n")) //2006/1/2 15:04:05一定要基于这个时间
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
