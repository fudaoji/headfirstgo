package section3

import (
	"fmt"
)

func Demo1() {
	persons := make(map[string]map[string]string, 2)
	persons["xm"] = map[string]string{"sex": "male"}
	persons["sb"] = map[string]string{"sex": "male"}

	fmt.Println(persons)
}
