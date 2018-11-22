package main

import "fmt"

// []int表示一个数组切片，[5]int表示含有5个元素的int数组
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	/*	(*arr)[0] = 100
		for i, v := range *arr {
			fmt.Println(i, v)
		}*/

	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 元素个数写在元素类型前面
	var arr1 [5]int
	arr2 := [4]int{1, 2, 3, 4}
	// 元素个数不能省略，可以用...表示自动获取
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	// 二维数组
	var grid [4][5]int
	fmt.Println(grid)

	// 遍历一维数组
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// 遍历二维数组
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}

	// 使用range关键字遍历一维数组
	for i := range arr2 {
		fmt.Println(arr2[i])
	}

	// 使用range关键字遍历二维数组
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}

	// 使用range关键字获取索引和元素
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 使用range关键字只获取数组中元素的值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr1)
	//printArray(arr2)// cannot use arr2 (type [4]int) as type [5]int in argument to printArray
	printArray(arr3)

	// 调用printArray函数之后arr[0]没有更新为100，说明数组是值传递
	fmt.Println(arr1, arr3)

	// [10]int 和 [20]int 是不同的类型
	// 调用 func f(arr [10]int) 会拷贝数组
	// go语言中一般不直接使用数组，一般使用切片

	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
