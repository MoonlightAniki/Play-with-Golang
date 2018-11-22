package main

import "fmt"

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

func main() {
	fmt.Println("Hello World");
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, ss)
}
