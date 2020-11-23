package main

import (
	"math"
	"math/cmplx"
)
import "fmt"

/*
内建变量类型
bool,string
(u（无符号）)int,(u)int8,(u)int32,,(u)int64,uintptr(指针，它的长度根据操作系统)
byte(8位), rune(32位，go语言的char)
float32,float64,complex64,complex128    complex是一种复数类型

类型转换是强制的，不能隐示转换类型
 */
func euler(){
	a := 3 +4i
	fmt.Println(cmplx.Abs(a))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) +1)
	fmt.Println(cmplx.Exp(1i*math.Pi) +1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi) +1)
}

func triangle(){
	var a,b int = 3,4
	var c int
	//c = math.Sqrt(a*a + b*b)   不能隐示转换类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	//euler()
	triangle()
}
