package main

import (
	"bufio"
	"errors"
	fib "fib/fib_i"
	"fmt"
	"os"
)

func tryDefer() {
	//defer 内部用的栈存的，所以先2后1,有了defer不怕异常return
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	/*
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	*/
	//错误处理
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error") //创建自定义error
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	f := fib.Fibonacci()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

//服务器统一错误处理

func main() {
	tryDefer()
	writeFile("fib.txt")
}
