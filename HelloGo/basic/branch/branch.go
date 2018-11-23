package main

import (
	"io/ioutil"
	"fmt"
)

// switch后可以没有表达式
// case子句默认会break，如果不要break需要使用fallthrough
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score : %d\n", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
		//fallthrough
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	/*	const filename = "abc.txt"
		// 函数可以有两个返回值
		contents, err := ioutil.ReadFile(filename)
		// if 语句的条件不需要加括号
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", contents)
		}
	*/

	// if 语句可以写成 for 语句的形式
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(-1), //panic会中断程序的执行
		grade(59),
		grade(60),
		grade(80),
		grade(99),
		grade(100),
		grade(101),
	)
}
