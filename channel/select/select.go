package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 1
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func work(i int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d \n", i, n)
	}

}

func CreateWork(i int) chan<- int {
	c := make(chan int)
	go work(i, c)
	return c
}

//func main() {
//	var c1, c2 = generator(), generator()
//	for {
//		select {
//		case n := <-c1:
//			fmt.Printf("recevied from c1 %d\n", n)
//		case n := <-c2:
//			fmt.Printf("recevied from c2 %d\n", n)
//		}
//	}
//}

//func main() {
//	var c1, c2 = generator(), generator()
//	var work = CreateWork(0)
//	var has_vaule = false
//	n := 0
//	for {
//		var activeWork chan <- int
//		if has_vaule{
//			activeWork = work
//		}
//		select {
//		case n = <-c1:
//			has_vaule = true
//		case n = <-c2:
//			has_vaule = true
//		case activeWork <-n:
//			has_vaule = false
//		}
//	}
//}

//这个运行是不会关闭的，需要让他关了
//func main() {
//	var c1, c2 = generator(), generator()
//	var work = CreateWork(0)
//	//var has_vaule = false
//	n := 0
//	var values []int
//
//	for {
//		var activeWork chan <- int
//		var activeValue int
//		if len(values) > 0{
//			activeWork = work
//			activeValue = values[0]
//		}
//		select {
//		case n = <-c1:
//			values = append(values, n)
//		case n = <-c2:
//			values = append(values, n)
//		case activeWork <- activeValue:
//			values = values[1:]
//		}
//	}
//}

func main() {
	var c1, c2 = generator(), generator()
	var work = CreateWork(0)
	//var has_vaule = false
	n := 0
	var values []int
	tm := time.After(5 * time.Second)
	tick := time.Tick(1 * time.Second)
	for {
		var activeWork chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWork = work
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWork <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Printf("bb")
			return
		case <-time.After(800 * time.Millisecond):
			fmt.Println("has no create")
		case <-tick:
			fmt.Println(len(values))
		}
	}
}
