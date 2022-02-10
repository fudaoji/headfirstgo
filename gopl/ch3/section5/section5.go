package section5

import (
	"fmt"
	"strconv"
)

func Demo() {
	num := 123

	println("fmt.Sprintf:", fmt.Sprintf("%d", num))
	println("strconv.Itoa:", strconv.Itoa(num))
	println("strconv.FormatInt:", strconv.FormatInt(int64(num), 10))
}
