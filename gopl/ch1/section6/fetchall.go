package section6

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//fetch  fetch content from url
func fetch(url string, ch chan string) {
	log.Println("======" + url + "开始======")
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%-8.2f%-15d%-200s", time.Since(start).Seconds(), nBytes, url)
	log.Println("======" + url + "结束======")
}

func FetchAll() {
	start := time.Now()
	ch := make(chan string)

	urls := os.Args[1:]

	for _, url := range urls {
		go fetch(url, ch)
	}

	fmt.Printf("%-8s%-15s%-200s\n", "time/s", "size/bit", "url")
	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Println("total time: ", time.Since(start).Seconds())
}
