package main

import "fmt"
/*
指针不能运算
值传递?引用传递
值传递是数据做了拷贝

go只有值传递这种方式
 */
func swap(a,b *int){
	*b,*a = *a,*b
	fmt.Println(*a, *b)
}

func swap2(a,b int){
	b,a = a,b
	fmt.Println(a, b)
}


func main() {
	//var a int = 2
	//var pa *int = &a
	//*pa = 6
	//fmt.Println(a)

	var a, b = 1 ,5
	var c, d = 3, 4
	swap(&a,&b)
	swap(&c,&d)
	fmt.Println(a,b, "xxxx")
	fmt.Println(c,d, "ddd")
}
