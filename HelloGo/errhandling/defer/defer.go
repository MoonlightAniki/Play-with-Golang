package main

import (
	"fmt"
	"bufio"
	"dev/Play-with-Go/HelloGo/functional/fib"
	"os"
	"errors"
)

func tryDefer() {
	/*	defer fmt.Println(1) //函数return之前调用，栈结构，后进先出
		defer fmt.Println(2)
		fmt.Println(3)
		//return
		panic("error occurred")
		fmt.Println(4)*/

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 自定义error
	err = errors.New("this is a custom error")

	if err != nil {
		//panic(err)
		//fmt.Println("Error:", err.Error())
		//fmt.Println("Error:", err)

		//Type Assertion
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}

		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
