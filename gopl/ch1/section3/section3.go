package section3

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func Exer1() {
	var counts map[string]int
	counts = make(map[string]int)
	files := make([]string, 0)
	//如果命令行参数有文件名则使用文件读取方式，否则使用命令行标准输入方式获取数据
	if args := os.Args; len(args) < 2 {
		countLines(os.Stdin, counts)
	} else {
		//1.打开文件：判断错误
		//2.defer 关闭文件
		//3.文件流放入函数
		for _, filename := range os.Args[1:] {
			counts = make(map[string]int)
			filepath := "./section3/" + filename
			file, err := os.Open(filepath)
			if err != nil {
				println("文件" + filepath + "不存在")
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				text := scanner.Text()
				if counts[text] >= 1 {
					files = append(files, filename)
					break
				} else {
					counts[text]++
				}
			}
		}
	}

	for _, filename := range files {
		println(filename)
	}
}

//Demo2  optimize file way
func Demo3() {
	counts := make(map[string]int)
	//如果命令行参数有文件名则使用文件读取方式，否则使用命令行标准输入方式获取数据
	if args := os.Args; len(args) < 2 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range os.Args[1:] {
			filepath := "./section3/" + filename

			//countLines(file, counts)
			content, err := ioutil.ReadFile(filepath)
			if err != nil {
				println(err)
				return
			}
			for _, line := range strings.Split(string(content), "\n") {
				counts[line]++
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			println(line, "=>", n)
		}
	}
}

//countLines  count line
func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}

//Demo2
func Demo2() {
	counts := make(map[string]int)
	//如果命令行参数有文件名则使用文件读取方式，否则使用命令行标准输入方式获取数据
	if args := os.Args; len(args) < 2 {
		countLines(os.Stdin, counts)
	} else {
		//1.打开文件：判断错误
		//2.defer 关闭文件
		//3.文件流放入函数
		for _, filename := range os.Args[1:] {
			filepath := "./section3/" + filename
			file, err := os.Open(filepath)
			if err != nil {
				println("文件" + filepath + "不存在")
			}
			countLines(file, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			println(line, "=>", n)
		}
	}
}

//Demo
func Demo() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, n := range counts {
		println(line, "=>", n)
	}
}
