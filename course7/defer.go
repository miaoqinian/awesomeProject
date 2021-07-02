package main

import (
	"awesomeProject/functionnal/fib"
	"bufio"
	"fmt"
	"os"
)

func testDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i * 2)
		if i == 20 {
			panic("error")
		}
	}
}

func openFileError(filename string) {
	_, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Err, pathError.Path, pathError.Timeout())
		}
	}
}

func main() {
	//writeFile("fib.txt")
	//tryDefer()
	openFileError("fib.txt")
}
