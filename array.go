package main

import "fmt"

var arr [10]int
var b = []float32{1.1, 3.4, 5.6}

func main() {
	arr[4] = 1
	fmt.Println(b)
	b[2] = 4.2
	fmt.Println(len(b))
	fmt.Println(arr[2:])

}
