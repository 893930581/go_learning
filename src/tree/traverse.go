package tree

func (node *TreeNode) Trverse() {
	if(node == nil){
		return
	}
	node.Left.Trverse()
	node.Print()//定义在下面了。
	node.Right.Trverse()
}