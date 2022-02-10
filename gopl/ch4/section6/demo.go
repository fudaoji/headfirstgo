package section6

import (
	"log"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func Demo1() {
	const text = `{{.Count}} items are made of {{.Material}}
`
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		log.Fatal(err)
	}
}
