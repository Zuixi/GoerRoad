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
	// 
} 