package main

import (
	"fmt"
	"math"
)

/*
常量
const
const()一组常量
const 数值可以当做各种类型来用
 */

func consts(){
	const filename = "abc.txt"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt(a*a +b*b))
	fmt.Println(c,filename)
}

func enums(){
	//const(
	//	cpp = 0
	//	java= 1
	//	python = 2
	//	golang = 3
	//)
	const(
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp,java,python,golang)
	// b,kb,mb,gb,tp,pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tp
		pb
	)
	fmt.Println(b,kb,mb,gb,tp,pb)
}

func main() {
	consts()
	enums()
}

/*
变量类型写在变量名后面
编译器可推测变量类型
没有char,只有rune
原生支持复数类型
 */