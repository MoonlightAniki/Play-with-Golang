package main

import (
	"dev/Play-with-Go/HelloGo/retriever/mock"
	"dev/Play-with-Go/HelloGo/retriever/real"
	"fmt"
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
	r = real.Retriever{}
	fmt.Println(download(r))
}
