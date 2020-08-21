package concurrency

import (
	"sync"
	"fmt"
)

// pool是对象池模式的并发安全实现
// 池模式是一种创建和提供固定数量的可用对象的方式,通常用于约束创建昂贵资源的事务(例如数据库连接)
// pool的主要接口是Get方法, Get首先检查pool中是否有可用实例返回给调用者,如果没有就新建一个变量
// 使用完成之后, 调用者使用Put方法将实例放回pool中,方便其他调用者使用

// -------------
// pool的简单示例
// -------------
func TestPool() {
	myPool := &sync.Pool{
		New: func() interface{} {
			// 如果创建实例,则打印相关信息
			fmt.Println("new instance was created")
			return struct{}{}
		},
	}

	// 这里调用Get方法, 由于实例尚未实例化,会调用New函数
	myPool.Get()
	// 这里也会调用New函数
	instance := myPool.Get()

	// 这里用完之后将其放回pool中,此时实例数量为1
	myPool.Put(instance)

	// 这里再次使用Get方法,但是New函数不会被调用,因为pool中含有一个实例
	myPool.Get()
}

// -----------------
// pool的使用场所
// -----------------

// pool可以预热分配对象的内存, 用于必须尽快执行的操作
// 在这种case下, 我们不是通过限制对象的数量来保护主机的内存, 而是通过预先加载获取另一个对象的引用来减少消费者的时间消耗
// 在编写网络服务器的时候, 这种情况非常常见

// 在TestPool中为什么不直接使用实例, 而是使用pool呢,因为在go的垃圾回收机制中, 任何实例化的对象都会被自动清理
// 可以参考下面的示例代码

func TestPoolWithCompare() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			// 返回字节切片的指针
			return &mem
		},
	}

	// 在这将池扩充到4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 << 10
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			// 断言此类型是一个指向字节切片的指针
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	// 如果这里不使用pool, 那么结果不是确定的
	fmt.Printf("%d calculators were created\n", numCalcsCreated)
}