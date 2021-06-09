package main

import "fmt"

func main() {
	var m = map[string]string{
		"name":   "monicx",
		"age":    "18",
		"gender": "男",
	}
	fmt.Println(m)
	m1 := map[string]int{}
	m1["monicx"] = 12
	fmt.Println(m1)

	var m2 = make(map[int]string)
	m2[1] = "monicx"
	fmt.Println(m2)

	delete(m, "name")

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("no key is name")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	var m3 = make(map[string]interface{})       //可以让值是多类型的值
	m3["age"] = 1
	m3["name"] = "mimi"


	var m4 = make(map [interface{}]interface{})
	m4[1] = "miao"
	m4["nian"] = 2.2
	for k, v := range m4 {
		fmt.Println(k, v)
	}

}
