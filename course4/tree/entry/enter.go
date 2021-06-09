package main

import (
	"awesomeProject/course4/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postLast() {
	if myNode == nil || myNode.node == nil{
		return
	}
	node := myTreeNode{myNode.node.Left}
	node.postLast()
	node2 := myTreeNode{myNode.node.Right}
	node2.postLast()
	myNode.node.Print()
}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = new(tree.Node)

	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	//root.SetValue(122)
	//root.Print()
	//var pRoot *tree.Node
	//pRoot.SetValue(200)
	root.Travel()
	//var pRoot *myTreeNode
	fmt.Println()
	var pRoot = myTreeNode{&root}
	pRoot.postLast()
}
