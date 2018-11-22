package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// 函数可以有多个返回值
// 带余除法
//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

// 可以给返回值取名，使代码可读性更好
//func div(a, b int) (quotient /*商*/ , remainder /*余数*/ int) {
//	return a / b, a % b
//}

// 第三种写法
// 但是不建议使用这种写法，方法中代码较多容易出错
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/*func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b) //不用的变量使用"_"代替
		return q
	default:
		panic("Unsupported operation: " + op)
	}
}*/

// 改写eval函数，更符合go语言风格
func eval(a, b int, op string) (res int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unsupported operation: %s\n", op)
	}
}

// 使用函数式编程实现eval函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum2(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	//fmt.Println(eval(10, 2, "x"))
	if res, err := eval(10, 2, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}

	//res, remainder := div(10, 3)
	//fmt.Println(res, remainder)

	fmt.Println(div(10, 3))

	fmt.Println(apply(pow, 3, 3))

	// 定义匿名函数
	fmt.Println(apply(func(a, b int) int {
		return a*a + b*b
	}, 3, 4))

	fmt.Println(sum2(1, 2, 3, 4, 5))
}
