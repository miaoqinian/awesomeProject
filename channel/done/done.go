package main

import (
	"fmt"
	"sync"
)

type work struct {
	in   chan int
	done func()
}

func createWork(i int, wg *sync.WaitGroup) work {
	w := work{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go dowrk(i, w)
	return w
}

func dowrk(i int, w work) {
	for n := range w.in {
		fmt.Printf("dodododo %d, %c\n", i, n)
		w.done()
	}

}

func chanDemo() {
	var wg sync.WaitGroup
	var works [10]work
	for i := 0; i < 10; i++ {
		works[i] = createWork(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		works[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		works[i].in <- 'f' + i
	}
	wg.Wait()

}

func main() {
	chanDemo()
}
