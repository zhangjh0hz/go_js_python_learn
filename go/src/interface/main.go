package main

import (
	"fmt"
	"interface/mock"
)

type Retiever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//接口组合
type RetieverPoster interface {
	Retiever
	Poster
}

const url = "http://www.imooc.com"

func download(r Retiever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(s RetieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked contents"})
	return s.Get(url)
}

func main() {

	var r Retiever
	rr := mock.Retriever{"this is fake imooc.com"}
	r = &rr
	if mockr, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockr.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Printf("%v\n", r)
	fmt.Println("Try a session")
	fmt.Println(session(&rr))
	//fmt.Println(download(r))

}
