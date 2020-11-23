package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func runForever(){
	for{
		fmt.Println("monicx is so handsome")
		time.Sleep(100)
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func runCoverToBin(n int) string {
	var result = ""
	for ;n >0;n/=2 {
		ml := n%2
		result = strconv.Itoa(ml) + result
	}
	return result
}

func main() {
	//runForever()
	readFile("a.txt")
	res := runCoverToBin(13)
	fmt.Println(res)
}
