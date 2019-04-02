package main

import (
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from "+"goroutine %d\n", i)
				a[i]++
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
}
