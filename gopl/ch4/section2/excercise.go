package section2

import (
	"fmt"
	"unicode"
)

func Q47() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(reverseByte(s)))
}

func Q46() {
	s := []rune{'h', 'e', ' ', ' ', 'l', 'l', ' ', ' ', 'o'}
	fmt.Println(string(removeRepeatSpace(s)))
}

func Q45() {
	s := []string{"he", "ll", "he", "he", "he"}
	fmt.Println(removeRepeatStr(s))
}

//reverseByte reverse byte slice
func reverseByte(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

//removeRepeatSpace  remove the repeate space in the slice
func removeRepeatSpace(str []rune) []rune {
	sLen := len(str)
	newStr := str[:0]
	for i := 0; i < sLen; i++ {
		if len(newStr) > 0 && unicode.IsSpace(str[i]) && unicode.IsSpace(newStr[len(newStr)-1]) {
			continue
		}
		newStr = append(newStr, str[i])
	}
	return newStr
}

//removeRepeatStr  remove the repeate ele in the slice
func removeRepeatStr(str []string) []string {
	sLen := len(str)
	newStr := str[:0]
	for i := 0; i < sLen; i++ {
		if len(newStr) > 0 && str[i] == newStr[len(newStr)-1] {
			continue
		}
		newStr = append(newStr, str[i])
	}
	return newStr
}
