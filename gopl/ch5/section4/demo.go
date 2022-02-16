package section4

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Demo1() {
	in := bufio.NewReader(os.Stdin)
	for {
		c, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("结束输入")
			break
		}

		if err != nil {
			fmt.Printf("Read fail: %v", err)
			return
		}

		fmt.Printf("%c", c)
	}

}
