package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"dev/Play-with-Go/HelloGo/tree"
)

// 一个目录中只能有一个包
// main()函数必须在main包下
// 为一个结构定义的函数必须放在同一个包内
func main() {
	var root tree.Node
	//fmt.Println(root.Val, root.Left, root.Right)
	fmt.Println(root)

	//     3
	//   /   \
	//  0     5
	//       / \
	//      4   2
	root = tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)
	fmt.Println(root)

	root.Print()
	fmt.Println()

	// 调用了setValue() 函数之后节点的值未改变
	// 使用指针版本之后可以改变节点的值
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	// 创建node切片
	/*	nodes := []Node{
			{3, nil, nil},
			{Val: 4},
			{5, nil, &root},
		}
		fmt.Println(nodes)*/

	var pRoot *tree.Node
	pRoot.SetValue(100) //nil可以调用函数
	pRoot = &tree.Node{Val: 200}
	pRoot.Print()
	fmt.Println()

	root.InOrderTraverse()

	fmt.Println()
	mRoot := myTreeNode{&root}
	mRoot.postOrderTraverse()

	fmt.Println()
	testSparse()

}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))

}

// 使用组合的方式对结构进行扩展
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrderTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}

	// cannot call pointer method on myTreeNode literal
	// cannot take the address of myTreeNode literal
	//myTreeNode{myNode.node.Left}.postOrderTraverse()
	//myTreeNode{myNode.node.Right}.postOrderTraverse()
	//myNode.node.Print()

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrderTraverse()
	right.postOrderTraverse()
	myNode.node.Print()
}
