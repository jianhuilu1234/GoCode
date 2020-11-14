package main

import "fmt"

type Phone interface {
	call()
}

type HWPhone struct {
}

func (hwPhone HWPhone) call() {
	fmt.Println("I am HW, I can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(HWPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
