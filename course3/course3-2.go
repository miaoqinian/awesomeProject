package main

import "fmt"

/*
slice
本身没有数据，是对底层array的一个view

slice是可以向后扩展的，不可以向前扩展
s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
 */
func main() {
	arr:= [...]int{0,1,2,3,4,5,6,7,8}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Println(s1) //[2 3 4 5]
	fmt.Println(s2) //[5 6]
	fmt.Println(s1[4:5]) //[5 6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s2,len(s2),cap(s2))
}
