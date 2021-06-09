package tree

func (node *Node) Travel(){
	if node == nil{
		return
	}
	node.Left.Travel()
	node.Print()
	node.Right.Travel()
}