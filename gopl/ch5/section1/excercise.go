package section1

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//Q51 excercise 5.1
func Q51() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "q51: %v", err)
	}

	links := visit1(nil, doc)
	fmt.Printf("len: %d\n", len(links))
	for _, link := range links {
		fmt.Println(link)
	}
}

//visit get links of a
func visit1(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit1(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit1(links, n.NextSibling)
	}
	return links
}
