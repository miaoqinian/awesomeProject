package main

import (
	"fmt"
	"time"
)

func work(i int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", i, n)
	}

}

func CreateWork(i int) chan<- int {
	c := make(chan int)
	go work(i, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWork(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func channelBuffer() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//channelClose()
	channelBuffer()
}
