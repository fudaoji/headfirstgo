package section4

import (
	"bytes"
	"os"
	"strings"
)

//Excer3 excersice 3.12 判断两个字符串是否同文异构
func Excer3() {
	args := os.Args
	str1, str2 := args[1], args[2]
	if homogeneous(str1, str2) {
		println("yes")
	} else {
		println("not")
	}
}

//homogeneous 同文异构判断
func homogeneous(str1, str2 string) bool {
	if str1 == str2 || len(str1) != len(str2) {
		return false
	}
	for _, c := range str1 {
		if !strings.Contains(str2, string(c)) {
			return false
		}
	}
	return true
}

//Excer2 excersice 3.11
func Excer2() {
	println(comma2("-121310.01"))
}

//comma2 adapt to float
func comma2(s string) string {
	//获取非符号或小数点的部分
	var pre string
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		pre = string(s[0])
		s = s[1:]
	}
	var suf string
	if i := strings.Index(s, "."); i >= 0 {
		suf = string(s[i:])
		s = s[:i]
	}
	return pre + comma(s) + suf
}

//Excer1 excersice 3.10
func Excer1() {
	println(comma1("-121310"))
}

//comma1
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buff bytes.Buffer
	mod := 3 - n%3
	for i, c := range s {
		buff.WriteRune(c)
		if (i+1+mod)%3 == 0 && i < n-1 {
			buff.WriteByte(',')
		}
	}
	return buff.String()
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func Demo() {
	println(comma("121311"))
}
