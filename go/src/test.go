package main

import (
	"fmt"
	"time"
)

func go_route(a int) {
	i := 0
	for {
		fmt.Print(i)
		i++
		time.Sleep(time.Second)
	}
}

type IRead interface {
	Reader()
}

type Student struct {
	Age     int
	Address string
	name    string
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, "ss"

	const (
		filename2 = "dbs.txt"
		c         = 3
	)
	fmt.Println("length = ", len(filename))
	fmt.Println(filename, a, b, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		kb = 1 << (10 * iota)
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(kb, mb, gb, tb, pb)
}

func main() {
	a := 1
	fmt.Println(a)
	consts()
	enums()
}
