package main

import (
	"awesomeProject/course5/my_test/mock"
	"awesomeProject/course5/my_test/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(p Poster) {
	p.Post("http://www.baidu.com",
		map[string]string{"name": "monicx",
			"course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
	OtherMethod(host string)
}

func main() {
	var r Retriever

	r = &real.Retriever{
		UserAgent: "chrome",
		TimeOut:   time.Minute,
	}
	switch r.(type) {
	case *real.Retriever:
		fmt.Println("this is true")

	}
	r = mock.Retriever{"www.baidu.com"}
	//type assertion
	if mockRetriver, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriver.Contents)
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
}
