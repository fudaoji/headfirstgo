package web

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type blogList struct {
	Total int
	List  []string
}

var filename string = "./web/blog.txt"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//reverseArr 数组逆置
func reverseArr(arr []string) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	t, err := template.ParseFiles("./web/index.html")
	check(err)

	var blogs []string
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		blogs = append(blogs, scanner.Text())
	}

	reverseArr(blogs)
	err = t.Execute(w, blogList{Total: len(blogs), List: blogs})

	/* tempText := "Hi, {{.}} this is a new template"
	t, err := template.New("test").Parse(tempText)
	check(err)
	err = t.Execute(os.Stdout, "Rocky") */

	check(err)
}

func newHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	t, err := template.ParseFiles("./web/new.html")
	check(err)
	err = t.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	content := req.FormValue("content")
	//save content
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(filename, options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, content)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(w, req, "/", http.StatusFound)
}

func Test8080() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
