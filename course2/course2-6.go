package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func apply(op func(int,int) int, a,b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling func %s with args (%d, %d)", opName,a,b)
	return op(a, b)
}

func pow(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

func manyArgs(numbers ...int){
	sum :=0
	for i := range numbers{
		sum += numbers[i]
	}
	fmt.Println(sum)
}

func main() {
	res := apply(pow,3,4)
	fmt.Println(res)
	manyArgs(1,2,3,4,5,6)

}
