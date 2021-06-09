package main

import (
	"awesomeProject/course4/queue"
	"fmt"
)

func main() {
	q := queue.Queue{3,4,5}
	fmt.Println(q.IsEmpty())
	q.Push(6)
	fmt.Println(q)
	head := q.Pop()
	fmt.Println(head)
	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println(q)
}
