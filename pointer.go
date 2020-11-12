package main

import "fmt"

func main() {
	var a int = 10
	b := &a
	var ptr *int

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*b)
	if ptr == nil {
		//println(ptr)
		fmt.Println(ptr)
		fmt.Printf("ptr 的值为 : %x\n", ptr)

	}

}
