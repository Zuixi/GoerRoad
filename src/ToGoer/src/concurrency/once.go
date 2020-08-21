package concurrency

import(
	"sync"
	"fmt"
)

// ----------
// Once的介绍
// ----------

// 定义
// type Once struct {
//    // contains filtered or unexported fields
// }
// Once对象只会执行一次
// func (o *Once)  Do(f func())
// Once的Do函数当且仅会执行一次里面的f函数

// Once的输出示例
func TestOnceWithChan() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only execute once")
	}
	done := make(chan bool)
	
	for i := 0; i < 10; i++ {
		// 在这里使用使用10个goroutine，但是只会执行一次
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// sync.Once只计算Do被调用的次数, 而不是调用传入Do的唯一函数一次
func TestOnceWithTwoFunc() {
	var count int
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Println("the res of call Do with different func is ", count)
}