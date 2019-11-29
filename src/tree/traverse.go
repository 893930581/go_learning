package tree

func (node *TreeNode) Trverse() {
	if(node == nil){
		return
	}
	node.Left.Trverse()
	node.Print()//定义在下面了。
	node.Right.Trverse()
}

func (node *TreeNode) Ftverse(f func(node *TreeNode)){
	if(node == nil){
		return
	}
	node.Left.Ftverse(f)
	f(node)
	node.Right.Ftverse(f)
}