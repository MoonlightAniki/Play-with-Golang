package main

import "fmt"

// Slice本身是没有数据的，是对底层数组的一个view(视图)
func updateSlice(s []int) {
	s[0] = 100
}

func printArray3(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// slice 可以向后扩展，不可以向前扩展
// s[i] 不可以超越 len(s), 向后扩展不可以超越底层数组cap(s)
func demo1() {
	fmt.Println("Extending slice:")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)
	s1 := arr[2:6 ]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s1)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 20)
	s5 := append(s4, 30)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println("arr =", arr)
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s := arr[2:6] //切片，左闭右开
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	// slice不是值类型，数组是值类型
	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("After updateSlice(s1):")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2):")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	arr2 := [5]int{1, 3, 5, 7, 9}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr2)
	fmt.Println(arr3)

	printArray3(arr2[:])
	printArray3(arr3[:])
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Reslice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	demo1()

}
