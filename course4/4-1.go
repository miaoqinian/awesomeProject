package main

import "fmt"

/*
go语言仅支持封装，不支持多态和继承（用接口来实现）
go没有class 只有struct
go没有构造函数这一说法。 可以用工产函数去实现
 */
type treeNode struct {
	value int
	left,right *treeNode
}

func (node treeNode) print()  {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int){
	//不加指针是值传递，是一种拷贝，加指针是地址传递，只有使用指针才能改变结构内容
	//nil 指针也可以调用方法。
	if node ==nil {
		fmt.Println("setting value to nil node,Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) travel(){
	if node == nil{
		return
	}
	node.left.travel()
	node.print()
	node.right.travel()
}


func createNode(value int) *treeNode{
	return &treeNode{value: value}  // 使用自定义工厂函数，注意返回了局剖变量的地址，也可以给外部用。
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.left.right = createNode(2)
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
	root.setValue(122)
	root.print()
	var pRoot *treeNode
	pRoot.setValue(200)

	root.travel()
}
