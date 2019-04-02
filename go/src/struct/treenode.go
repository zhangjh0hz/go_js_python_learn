package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(v int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored.")
		return
	}
	node.value = v
}

func (node *treeNode) traverse() {
	/*
		if node == nil {
			return
		}
		node.left.traverse()
		node.print()
		node.right.traverse()
	*/
	node.TraverseFunc(func(n *treeNode) {
		n.print()
	})
	fmt.Println()
}

func (node *treeNode) TraverseFunc(f func(*treeNode)) {
	if node == nil {
		return
	}

	node.left.traverse()
	f(node)
	node.right.traverse()
}

//扩展已有类型，
type MyTreeNode struct {
	node *treeNode
}

func (mynode *MyTreeNode) postOrder() {
	if mynode == nil || mynode.node == nil {
		return
	}

	left := MyTreeNode{mynode.node.left}
	left.postOrder()
	right := MyTreeNode{mynode.node.right}
	right.postOrder()
	mynode.node.print()
}

func main() {
	//创建
	var root treeNode
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	fmt.Println(root)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	root.left.right.setValue(11)
	root.left.right.print()
	fmt.Println()

	root.traverse()

	nCount := 0
	root.TraverseFunc(func(n *treeNode) {
		nCount++
	})

	fmt.Println("count = ", nCount)
}
