package concurrency

import (
	"fmt"
)

func sayhello() {
	fmt.Println("Hello")
}

func Test() {
	go sayhello()
	fmt.Println("Test")
}


func TestAnonmous() {

	// 匿名函数前需要紧跟func
	go func() {
		fmt.Println("Say Anonmous go routine")
	}()

	fmt.Println("Anonmous")
}

func TestAnonmousPlus() {
	
	Sayhello := func() {
		fmt.Println("use anonmous function")
	}
	go Sayhello()

	fmt.Println("anonmous plus")
}
