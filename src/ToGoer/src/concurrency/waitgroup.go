package concurrency

import (
	"fmt"
	"sync"
)

func CreateJoinKnot() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()    // 这里将加入连接点，确保可以阻塞main goroutine
}

func TestJoinScope() {
	var wg sync.WaitGroup
	str := "Out"

	wg.Add(1)
	// 测试闭包的词组范围
	func() {
		defer wg.Done()
		str = "Inner"		
	}()
	wg.Wait()    // 这里将加入连接点，确保可以阻塞main goroutine
	
	// 这里将会输出Inner
	// goroutine在它创建的同一地址空间内执行
	fmt.Println("str is ", str)
}

func TestJoinScopePro() {
	var wg sync.WaitGroup
	
	for _, val := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("str is ", val)
		}()
	}

	// 这里的输出结果和我们想象的不一样，我这是good day good day greetings
	wg.Wait()
}

func TestJoinScopeProModify() {
	var wg sync.WaitGroup
	
	// 为了每个goroutine可以输出自己的string，需要传入对应的副本
	for _, val := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			fmt.Println("str is ", val)
		}(val)
	}

	// 这里的输出结果和我们想象的一样
	wg.Wait()
}