package main

import "fmt"

//数组 作为参数传递时，是值传递，会拷贝一份
func printArray(arr [3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//需要更改的话，一种方法用指针
func printArrayPointer(arr *[3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//另外一种方法是用slice
func updateWithSlice(arr []int) {
	arr[0] = 1000
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := []int{2, 3, 4, 5}
	arr4 := [...]int{2, 3, 4, 5}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, arr4, grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	printArray(arr2)
	fmt.Println(arr2)
	printArrayPointer(&arr2)
	fmt.Println(arr2)

	updateWithSlice(arr4[:]) //传参用[:]表示切片
	fmt.Println(arr4)
}
