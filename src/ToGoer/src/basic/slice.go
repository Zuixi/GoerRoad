package basic

import (
	"fmt"
	_"unicode/utf8"
)

// 引用类型 ：slice, map, channel, function
// 引用类型的零值是 nil
func TestSlice() {
	// ---------
	// 声明和初始化
	// ---------

	// 创建一个含有4个元素的slice
	// make是一个特殊的内置函数，仅适用于slce、map、channel
	// 在下面的例子中，make创建一个含有4个string的slice，创建完成之后可以得到对应的数据结构如下：
	// 第一个是返回数组的指针，第二个是的长度，第三个是slice的容量
	// --------
	// | * | --> | nil | | nil | | nil | | nil |
	// -------   | 0   | | 0   | | 0   | | 0   |
	// | 4 |
	// -------
	// | 4 |
	// -------

	// ---------------
	// 长度和容量的区别
	// ---------------

	// 长度是指针指向的我们可以访问的元素的个数
	// 容量是指针指向的数组中总的元素的个数

	// 语法糖 --> 看起来很像数组
	// 需要注意的一件事情是：make创建的slice中是没有值的，这和数组有很大的区别
	slicestring := make([]string, 4)
	slicestring[0] = "Apple"
	slicestring[1] = "Banana"
	slicestring[2] = "Grape"
	slicestring[3] = "Plum"

	// 不能访问超过slice长度的元素，或者会报错， index out of range
	// slicestring[5] = "Runtime error"

	fmt.Println("slicestring is ", slicestring)
	for _, val := range(slicestring) {
		fmt.Println("slicestring is ", val)
	}
	fmt.Println()
	
	// --------------------
	// 使用引用类型创建slice
	// --------------------

	// 创建一个长度为5，容量为8的slice
	// make可以指定创建slice的容量和长度在初始化的时候
	// ----------
	// | * | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	// ----		 |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// | 5 |
	// | 8 |
	// ----------
	// 表示可以读取前5个元素并且还有三个元素的容量，这三个元素可以之后再用
	slice2 := make([]string, 5, 8)
	slice2[0] = "A"
	slice2[1] = "B"
	slice2[2] = "C"
	slice2[3] = "D"
	slice2[4] = "E"

	fmt.Println("capacity 和 length的区别-->")
	inspectSlice := func(slice []string) {
		fmt.Printf("Length[%d], Capacity[%d]\n", len(slice), cap(slice))
		for i := range slice {
			fmt.Printf("[%d] %p %s \n",i , &slice[i], slice[i])
		}
	}

	inspectSlice(slice2)

	//-------------------------------------------------
	// 更进一步， 让slice成为动态的数据结构，就和vector一样
	//-------------------------------------------------

	// 声明一个string的slice，一开始将是0值
	// 结构分别是nil, 0, 0, 使用var 声明将会是nil
	var name []string

	// name := string{}, 这和name slice是不一样的， 因为在这里name将不会赋予nil类型
	// var会自动为变量赋予零值，而其他创建的空值类型的变量不一定总是会自动赋予零值，这是区别
	// 上面声明的name slice含有一个指针，指向nil, 所以这是一个空的slice,而不是一个nil slice
	// 空slice和nil slice之间还是有区别的，任何引用类型的值都可以被设为零值nil，

	// 获取slice的capacity
	lastCapacity := cap(name)
	fmt.Printf("last cap of name is [%d]\n", lastCapacity)

	// 往name中添加10000个元素
	for index := 1; index <= 10000; index++ {
		name = append(name, fmt.Sprintf("Name: %d", index))

		// 每次使用append方法时候，都会检查slice的容量和长度
		// 如果获取的长度和容量是一样的，意味着slice没有空间存放元素了；这个时候需要返回一个新建的数组，容量是之前的两倍
		// 复制原来的元素后在添加新的元素， 在go的堆栈上slice副本发生变化，并且会返回新的引用，我们需要用这个新的引用代替之前的
		// 如果获取的数据是不一样的，那么意味着还有空间可以存放元素，这个时候效率比较高

		// 注意下面输出的最后一列，当返回的数组元素小于或等于1000个元素的时候，增长速度是一倍，超过这个数字的时候，增长速度变成25%
		if lastCapacity != cap(name) {
			// 计算变速速率
			capChg := float64(cap(name) - lastCapacity) / float64(lastCapacity) * 100

			// 保存新值
			lastCapacity = cap(name)

			// 输出对应的结果
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &name[0], index, cap(name), capChg)
		}
	}


	// ----------------
	// slice of slice
	// ----------------

} 