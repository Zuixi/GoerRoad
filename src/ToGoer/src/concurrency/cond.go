package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// ----------
// Cond 的定义形式
// ----------

// type Cond struct {
//     // L is held while observing or changing the condition
//     L Locker
//     // contains filtered or unexported fields
// }
// -------------
// for 的粗暴使用
// -------------
// Cond实现了一个条件变量，用于等待或者宣布事件发生时goroutine的交汇点
// 事件是指两个或者多个goroutine之间的任何信号
// 如果想要在收到某个goroutine信号之前让其处于等待状态，不是用Cond一种粗暴方法是
// for conditionTrue() == false {
// }	
// 循环等待执行效率太低，可以在循环体中加入sleep改善，但是仍旧效率低下

// -------------
// Cond 的使用
// -------------
// cond := sync.NewCond(&sync.Mutex{}) 
// cond.L.Lock()
// for conditionTrue == false{
// 	cond.Wait()
// }
// cond.L.Unlock()

// 上面wait的调用不仅仅是阻塞，暂停当前的goroutine且允许其他的goroutine在机器上运行
// 还会发生一些其他事情-->
// 进入wait之后，Cond的Locker变量会调用UnLock并且在退出wait的时候使用Lock
// 看起来在等待条件发生的整个过程中都拥有这个锁，实际上并不是

// ---------
// 具体例子
// ---------

// 一个等待信号的goroutine 和 发送信号的goroutine
// 假设有一个固定长度为2的队列，并且需要将10个元素放入队列中
// 希望做到队列中一有空间就可以放入元素，所有队列在有空间时需要立刻发送信号通知

func TestCond() {
	// 使用sync.Mutex来创建Cond
	cond := sync.NewCond(&sync.Mutex{})

	// 声明一个初始长度0，容量10的队列
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// 再次进入该并发条件下的关键部分，以修改与并发条件判断直接相关的数据
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		// 退出操作关键部分，因为我们已经成功移除了一个元素
		cond.L.Unlock()
		// 发出信号，通知处于等待状态的goroutine可以进行下一步
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		// 在发送信号之前Lock L
		cond.L.Lock()

		// 检查队列长度，用于确认什么时候需要等待
		// removeFromQueue 是异步执行，if 没法重复判断， for才行
		// 在这里使用if 语句会出现错误的
		for len(queue) == 2 {
			// 使用wait 阻塞main goroutine
			cond.Wait()
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})

		// 新建goroutine, 将会在1s后移除第一个元素
		go removeFromQueue(1 * time.Second)
		// 成功加入一个元素，需要退出goroutine
		cond.L.Unlock()
	}

	fmt.Println("\n----Test Cond Over----\n")
}