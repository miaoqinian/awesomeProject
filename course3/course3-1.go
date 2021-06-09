package main

import "fmt"
/*
数组是值类型
 */

func printArray(arr *[5]int){
	arr[0]=99
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}

func printArray2(arr []int){
	arr[0]=99
	//for i,v :=range arr{
	//	fmt.Println(i,v)
	//}
}


func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1,2,3,4,5}   // ...是让编译器自己数数组有多少个
	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//遍历数组的方法
	//for i:=0; i<len(arr3);i++{
	//	fmt.Println(arr3[i])
	//}

	//for index, value := range arr3{
	//	fmt.Println(index,value)
	//}

	//for index := range arr3{
	//	fmt.Println(index)
	//}

	printArray(&arr3)
	//printArray2(arr3[:])
	fmt.Println(arr3)



}
