package section1

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func DemoOutline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

//outline print html node
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for i := n.FirstChild; i != nil; i = i.NextSibling {
		outline(stack, i)
	}
}
