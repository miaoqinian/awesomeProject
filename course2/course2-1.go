package main
/*
变量定义
使用var关建字
var a,b,c bool
var s1,s2 string = "hello","world"
可以放在函数内，或直接放在包内
使用var()集中定义变量

让编译器自动决定类型
var a,b,i,s1,s2 = true,false,3,"hello","world"
使用：=定义变量  只能在函数内使用。
 */
import "fmt"

func variableZero(){
	var a int
	var s string
	fmt.Println(a, s, a, a)
}
var (
	monicx = "缪起年"
	andy = '?'
)

func variableMany(){
	cc, dd, ee := "cc", 8, '7'
	println(cc, dd, ee)
}

func main() {
	//fmt.Println("hello world")
	variableZero()
	println(monicx, andy)
	variableMany()
}
