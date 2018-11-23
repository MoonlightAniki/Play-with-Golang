package tree

import "fmt"

// go语言支持封装，不支持继承和多态
// 首字母大写表示public

type Node struct {
	Val         int
	Left, Right *Node
}

// 定义成员函数
func (node Node) Print() {
	fmt.Print(node.Val, " ")
}

// 等价于
/*func Print(node Node) {
	fmt.Print(node.Val)
}*/

// 值传递
/*func (node Node) SetValue(Val int) {
	node.Val = Val;
}*/

// 为了能改变节点的值，可以使用指针
func (node *Node) SetValue(val int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Val = val; //go语言中的指针可以使用.获取成员
}

// go语言没有构造函数，使用工厂函数代替
func CreateNode(val int) *Node {
	return &Node{Val: val} //返回一个局部变量的地址？在C++中是不合法的
}
