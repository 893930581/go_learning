package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//go语言结构体没有this指针，所以要给结构体定义方法就需要定义结构接收者方法。

func createNode(value int) *treeNode {
	return &treeNode{value: value}
	//go语言可以返回局部变量地址。
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

//但是其实接收者方法和普通方法是一样的，只是可以规定调用者
//接受者非常灵活，不管是定义的是指针接受者还是值接收者
//传入时既可以是指针又可以是值
func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{5, nil, nil}
	root.right = &treeNode{}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.left.right = createNode(2)
	root.left.right.print()
	root.left.right.print()
	root.left.right.print()
}
