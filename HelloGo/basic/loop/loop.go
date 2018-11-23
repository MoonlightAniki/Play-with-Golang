package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
	"io"
)

func sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// 将十进制数组转成二进制字符串
func convertToBinaryString(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

// 逐行打印文件中的文本
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// for 语句无循环初始条件和循环递增条件的话可以省略分号
	// 等价于java语言中的while语句
	// go语言中没有while关键字
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	// true 可以省略
	for {
		fmt.Println("abc")
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(sum(100))
	fmt.Println(
		convertToBinaryString(10),
		convertToBinaryString(13),
		convertToBinaryString(32),
		convertToBinaryString(23333),
		convertToBinaryString(0),
	)

	printFile("basic/abc.txt")

	s := `hello
	world
	kkkk
	dddd

	p`

	printFileContents(strings.NewReader(s))

	//forever()
}
