package section1

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c ByteCounter) String() string {
	return fmt.Sprintf("len: %d", c)
}

func DemoByteCounter() {
	var n ByteCounter
	var bs = make([]byte, 0)
	bs = append(bs, 'h')
	bs = append(bs, 'e')
	bs = append(bs, 'l')
	n.Write(bs)

	fmt.Println(n)

	n = 0
	var name = "world"
	fmt.Fprintf(&n, "hello, %s", name)
	fmt.Println(n)
}
