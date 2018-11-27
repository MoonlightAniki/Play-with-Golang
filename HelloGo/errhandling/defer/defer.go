package main

import (
	"fmt"
	"os"
	"bufio"
	"dev/Play-with-Go/HelloGo/functional/fib"
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
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
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
	tryDefer()
	//writeFile("fib.txt")
}
