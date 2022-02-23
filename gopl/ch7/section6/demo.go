package section6

import (
	"fmt"
	"reflect"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var slice []string = []string{"zacky", "rocky", "doogie"}
var intArr []int = []int{1, 10, 3}
var intArr1 []int = []int{1, 10, 3, 0}

//DemoDeepEqual 比较两个slice
func DemoDeepEqual() {
	fmt.Println(reflect.DeepEqual(intArr, intArr1))
}

//DemoIntSort 使用内置方法排序整型slice
func DemoIntSort() {
	sort.Ints(intArr)
	fmt.Println(intArr)
}

//DemoStringSort 使用内置方法排序字符串slice
func DemoStringSort() {
	sort.Strings(slice)
	fmt.Println(slice)
}

//DemoMySort 自定义实现字符串slice排序
func DemoMySort() {
	sort.Sort(StringSlice(slice))
	fmt.Println(slice)
}
