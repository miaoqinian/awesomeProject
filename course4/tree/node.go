package tree

import "fmt"

/*
go语言仅支持封装，不支持多态和继承（用接口来实现）
go没有class 只有struct
go没有构造函数这一说法。 可以用工产函数去实现
 */
type Node struct {
	Value int
	Left,Right *Node
}

func (node Node) Print()  {
	fmt.Print(node.Value,"　")
}

func (node *Node) SetValue(value int){
	//不加指针是值传递，是一种拷贝，加指针是地址传递，只有使用指针才能改变结构内容
	//nil 指针也可以调用方法。
	if node ==nil {
		fmt.Println("setting Value to nil node,Ignored.")
		return
	}
	node.Value = value
}


func CreateNode(value int) *Node {
	return &Node{Value: value} // 使用自定义工厂函数，注意返回了局剖变量的地址，也可以给外部用。
}

