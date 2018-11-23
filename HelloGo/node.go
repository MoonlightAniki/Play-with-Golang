package main

import "fmt"

// go语言支持封装，不支持继承和多态

type treeNode struct {
	val         int
	left, right *treeNode
}

// 定义成员函数
func (node treeNode) print() {
	fmt.Print(node.val)
}

// 等价于
/*func print(node treeNode) {
	fmt.Print(node.val)
}*/

// 值传递
/*func (node treeNode) setValue(val int) {
	node.val = val;
}*/

// 为了能改变节点的值，可以使用指针
func (node *treeNode) setValue(val int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.val = val; //go语言中的指针可以使用.获取成员
}

// 二叉树的中序遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Print(node.val, " ")
	node.right.traverse()
}

// go语言没有构造函数，使用工厂函数代替
func createNode(val int) *treeNode {
	return &treeNode{val: val} //返回一个局部变量的地址？在C++中是不合法的
}

func main() {
	var root treeNode
	//fmt.Println(root.val, root.left, root.right)
	fmt.Println(root)

	root = treeNode{val: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.right.right = createNode(2)
	fmt.Println(root)

	root.print()
	fmt.Println()

	// 调用了setValue() 函数之后节点的值未改变
	// 使用指针版本之后可以改变节点的值
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	// 创建node切片
	/*	nodes := []treeNode{
			{3, nil, nil},
			{val: 4},
			{5, nil, &root},
		}
		fmt.Println(nodes)*/

	var pRoot *treeNode
	pRoot.setValue(100) //nil可以调用函数
	pRoot = &treeNode{val: 200}
	pRoot.print()
	fmt.Println()

	root.traverse()

}
