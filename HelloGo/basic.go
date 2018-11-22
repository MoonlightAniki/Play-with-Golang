package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

//var aa = 3
//var bb = false
//var ss = "guoliang"

// 方法外定义的变量是包内私有的，多个变量可用 var() 定义
var (
	aa = 3
	bb = false
	ss = "guoliang"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //%q表示将变量用引号包括打印
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	//var a = 3
	//var s = "abc"
	//fmt.Println(a, s)

	var a, b, s = 1, true, "abc"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, s := 2, false, "bcd"
	fmt.Println(a, b, s)
}

// 验证欧拉公式
// pow(e, i * PI) + 1 = 0
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 类型转换是强制的

// 没有char类型，但是有rune类型
func triangle() {
	var a int = 3
	var b int = 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	// c = math.Sqrt(a*a + b*b) 是错误的，int类型不能自动转成float64类型，float64类型也不能自动转成int类型，必须强制类型转换
	fmt.Println(c)
}

// 常量的定义
// const数值可作为各种类型使用
func consts() {
	const filename string = "abc.txt"
	// 如果声明了常量的类型，常量的类型就确定了，转成其他类型需要强制类型转换
	//const a, b int = 3, 4
	// 如果未声明常量类型，常量类型就不确定，可以自动类型转换
	const a, b = 3, 4

	// 常量也可以定义在一个组内
	//const (
	//	filename string = "abc.txt"
	//	a, b            = 3, 4
	//)

	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

// 使用常量定义枚举类型
func enums() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)

	//const (
	//	cpp = iota//表示从0自增
	//	java
	//	python
	//	golang
	//)

	const (
		cpp        = iota
		_           // 不需要的可用"_"跳过
		python
		golang
		javascript
	)

	fmt.Println(cpp, javascript, python, golang)

	// iota还可以参与运算
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World");
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, ss)

	euler()

	triangle()

	consts()

	enums()
}
