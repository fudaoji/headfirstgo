package section5

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Excer3  excersice 3  输出状态码 resp.Status
func Excer3() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", content)
		fmt.Println("状态：", resp.Status)
	}
}

//Excer2  excersice 2  增加http//前缀，使用strings.HasPrefix判断
func Excer2() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", content)
	}
}

//Excer1  excersice 1
func Excer1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		//var contentW io.Writer
		res, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n==============%d=================", res)
	}
}

func Demo() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", content)
	}
}
