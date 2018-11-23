package main

import (
	"dev/Play-with-Go/HelloGo/retriever/mock"
	"dev/Play-with-Go/HelloGo/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	//var r Retriever
	//fmt.Println(download(r))

	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	// 类型声明失败
	// panic: interface conversion: main.Retriever is *real.Retriever, not mock.Retriever
	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Contents)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock Retriever")
	}

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Printf("UserAgent:%s, TimeOut:%d\n", v.UserAgent, v.TimeOut)
	}
}
