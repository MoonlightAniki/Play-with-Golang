package tree

import "fmt"

// 将二叉树的遍历函数都放在一个文件中

// 二叉树的中序遍历
func (node *Node) InOrderTraverse() {
	node.InOrderTraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 前序遍历
func (node *Node) PreOrderTraverse() {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	node.Left.PreOrderTraverse()
	node.Right.PreOrderTraverse()
}

// 后序遍历
func (node *Node) PostOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.PostOrderTraverse()
	node.Right.PostOrderTraverse()
	fmt.Print(node.Val, " ")
}

// 函数式编程实现中序遍历
func (node *Node) InOrderTraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.InOrderTraverseFunc(f)
	f(node)
	node.Right.InOrderTraverseFunc(f)
}
