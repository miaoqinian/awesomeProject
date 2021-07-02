package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				fmt.Printf("balabala %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
