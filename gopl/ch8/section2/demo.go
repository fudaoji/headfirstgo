package section2

import (
	"io"
	"log"
	"net"
	"time"
)

//DemoGoroutine 使用goroutine演示fib
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
