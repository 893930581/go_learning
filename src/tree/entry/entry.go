package main
import ".."
import "fmt"

type myTreeNode struct{
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrderTrverse (){
	if myNode == nil || myNode.node == nil {
		return 		
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrderTrverse()
	right.postOrderTrverse()
	myNode.node.Print()
}

func main() {
	// var groot struct{}
	var root tree.TreeNode
	// 在go中不论是实例还是指针都使用.来调用
	root = tree.TreeNode{Value: 3}
	// 创建方法一
	root.Left = &tree.TreeNode{Value:5, Left:nil, Right:nil}
	// 创建方法二
	root.Right = &tree.TreeNode{}
	// 创建方法三
	root.Right.Left = new(tree.TreeNode)
	// 创建方法四
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{Value:6, Left:nil, Right:&root},
	}
	// 创建方法五
	root.Left.Right = tree.CreateNode(2)
	// 工厂式创建
	fmt.Println(root,root.Left,root.Right,root.Left.Right,nodes)

	myRoot := myTreeNode{&root}
	myRoot.postOrderTrverse()
}

