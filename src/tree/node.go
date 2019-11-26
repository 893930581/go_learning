package tree

import "fmt"

type TreeNode struct {
	//声明结构体
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) /**接收者*/ Print() {
	//定义接收者函数,相当于其他语言对象中的this
	//JS Person.prototype.print = function(){}
	fmt.Println(node.Value)
}
//go语言结构体没有this指针，所以要给结构体定义方法就需要定义结构接收者方法。

func (node *TreeNode) SetValue(Value int) {
	node.Value = Value
}
//但是其实接收者方法和普通方法是一样的，只是可以规定调用者
//接受者非常灵活，不管是定义的是指针接受者还是值接收者
//传入时既可以是指针又可以是值

func CreateNode(value int) *TreeNode /**返回一个指针类型*/ {
	//自定义工厂函数,类似js中的constructor
	return &TreeNode{Value: value}
	//go语言可以返回局部变量地址,如果调用之后返回的局部变量的地址被别人使用
	//就不会被垃圾回收,否则反之
}