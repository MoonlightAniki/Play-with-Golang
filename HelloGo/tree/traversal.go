package tree

import "fmt"

// 将二叉树的遍历函数都放在一个文件中

// 二叉树的中序遍历
func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.InOrderTraverse()
	fmt.Print(node.Val, " ")
	node.Right.InOrderTraverse()
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
