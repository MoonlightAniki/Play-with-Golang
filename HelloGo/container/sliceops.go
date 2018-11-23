package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating Slice:")

	var s []int // Zero value for slice is nil
	fmt.Println(s == nil)

	for i := 1; i <= 100; i++ {
		fmt.Printf("i=%d, len(s)=%d, cap(s)=%d\n", i, len(s), cap(s))
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// slice的初始化
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)

	// 创建一个长度为16的slice
	s2 := make([]int, 16) //16是len
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	// slice的拷贝
	fmt.Println("Copying Slice:")
	copy(s2, s1)
	printSlice(s2)

	// 删除slice中的元素
	fmt.Println("Deleting elements from slice:")
	s2 = append(s2[:3], s2[4:]... /*可变参数*/) //删除s2中下标为3的元素
	printSlice(s2)

	// 从slice的头尾删除元素
	fmt.Println("Popping from front:")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front =", front)
	printSlice(s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail =", tail)
	printSlice(s2)

}
