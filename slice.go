package main

import "fmt"

func main() {
	numbers := []int{1, 3, 5, 6, 9, 110}
	printSlice(numbers)

	numbers1 := make([]int, 5, 9)
	printSlice(numbers1)
	numbers1 = append(numbers1, 1, 3, 4)
	printSlice(numbers1)

	//	创建切片是之前切片的两倍
	numbers2 := make([]int, len(numbers), cap(numbers)*2)

	//拷贝numbers到numbers2
	copy(numbers2, numbers)
	printSlice(numbers2)
}

func printSlice(x []int) {
	fmt.Println("len=%d cap=%d slice=%v", len(x), cap(x), x)
}
