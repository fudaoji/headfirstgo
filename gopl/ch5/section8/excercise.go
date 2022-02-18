package section8

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func Q518() {
	url := os.Args[1]
	path, n, err := fetch1(url)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("fetch %s to file %s, len: %d", url, path, n)
}

func fetch1(url string) (filename string, n int64, err error) {
	//打开网页
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	//保存内容到页面
	path := path.Base(resp.Request.URL.Path)
	if path == "/" || path == "." {
		path = "./section8/q518.html"
	}

	f, err := os.Create(path)
	if err != nil {
		return "", 0, err
	}
	defer func() { err = f.Close() }()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}
	return path, n, err
}
