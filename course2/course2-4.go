package main

import (
	"fmt"
	"io/ioutil"
)

/*
if的条件里可以赋值
if的条件里赋值的变量作用域就在这个if语句里

switch会自动break,除非使用fallthrough
 */

func scoreLevel(score int,age int) string {
	level := ""
	switch age < 10 {
	case score<0 || score > 100:
		panic(fmt.Sprintf("the score is wrong %d", score))
	case score < 60:
		level = "D"
		fallthrough
	case score < 70:
		level = "C"
	case score < 80:
		level = "B"
	case score < 90:
		level = "A"
	case score < 100:
		level = "S"
	}
	return level
}

func main() {
	const filename = "a.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err == nil {
	//	fmt.Println(contents)
	//}else {
	//	fmt.Println(err)
	//}
	if contents, err := ioutil.ReadFile(filename); err == nil {
		//fmt.Printf("%s\n", contents)
		fmt.Println(string(contents))
	}else {
		fmt.Println(err)
	}

	//scoreLevel(-1)
	res := scoreLevel(50,5)
	fmt.Println(res)
}
