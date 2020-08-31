package concurrency

import (
	_"sync"
	"fmt"
)

// 使用并发代码时，安全操作有以下几种不同的操作
// 用于共享内存的同步原语(sync.Mutex)
// 通过通信(channel)同步
// 不可变数据， 隐式安全
// 受限制条件保护的数据，隐式安全

// 不可变数据是最理想的，因为在并发程序中不会改变这个数据
// 不可变数据的使用依赖于约定, 坚持约定很难在项目上协调

func ConstrainsData() {
	// make slice 
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	// 从channel中取出对应值
	for num := range handleData {
		fmt.Printf("%d \t", num)
	}
}

// 在上面的例子中 loopData好handleData channel都是用了slice data, 但只有loopData对slice进行了直接访问
// 这里没有对slice做显式的访问和操作约束，有可能会产生问题
// 词汇约束，使用编译器对对变量的操作进行约束
// 词法约束涉及使用词法作用域仅公开用于多个并发进程的正确数据和并发原语




 