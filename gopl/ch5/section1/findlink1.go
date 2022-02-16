package section1

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//visit get links of a
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func DemoFindlinks1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v", err)
	}

	links := visit(nil, doc)
	fmt.Printf("len: %d\n", len(links))
	for _, link := range links {
		fmt.Println(link)
	}
}
