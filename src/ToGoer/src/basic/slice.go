package basic

import (
	"fmt"
	"unicode/utf8"
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

	// 在这我们将name切片取出一部分当做另外一个slice
	// namecopy 切片的长度是2
	// 参数是[开始索引: 开始索引 + 长度]，左闭右开
	// 通过看输出结果，可以知道这两个切片是在共享内存
	namecopy := name[2: 4]
	fmt.Printf("\n=> Slice of slice (before)\n")

	inspectSlice(name)
	inspectSlice(namecopy)

	// 既然是引用，那么如果我们改变namecopy[0]的值，那么name[0]应该也会变化
	namecopy[0] = "Modify Value"

	// 所以如果需要修改某些切片的值，那么我们就必须时刻知道有谁在用这个切片
	fmt.Printf("\n=> Slice of slice (after)\n")
	inspectSlice(name)
	inspectSlice(namecopy)

	// namecopy := append(namecopy, "CHANGED") 的行为是怎么样的
	// 如果长度和容量不一样的话，那么append中会有同样的问题产生
	// 由于namecopy的长度是2，容量是6，所以还有其他空间可以容纳元素
	// 但是这样将会改变name切片的元素，这不符合常理

	// -----------
	// 复制一个切片
	// -----------

	// copy函数只对string和切片有效
	// 声明一个容量足够的slice，让他能够存储slice2的全部元素
	// copy 函数的使用

	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)
	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	// -------------
	// 切片和引用类型
	// -------------

	// 声明一个包含5个整数的切片
	x := make([]int, 5)

	// 随便赋予初值
	for i := 0; i < 5; i++ {
		x[i] = i * 100
	}

	// 将x切片的第二个元素的地址赋予指针
	pointer := &x[1]

	// 往X中新加一个元素1000， 这将会触发错误
	// 由于x中长度和容量一样都是5，那么append方法将会扩容一倍
	// 所以pointer指针将会指向不一样的内存区域
	x = append(x, 1000)

	// 当我们改变第二个元素的值，pointer指针将不会发生变化，因为指针指向的是旧的x切片
	// 所以每次使用pointer的值的时候，都会读取到一个错误的值
	x[1]++

	// 通过查看输出结果，我们就可以看出
	fmt.Printf("\n=> 切片和引用类型\n")
	fmt.Println("Pointer: ", *pointer, "\t x[1:] ", x[1])

	// ----------------
	// UTF-8
	// ----------------
	fmt.Printf("\n=> UTF-8\n")

	// go中的字符集都是基于utf-8的
	// 所以当使用不同的编码的时候，有可能会引发一些其他问题

	// 声明一个字符串，里面包含中文和英语字符
	// 对于每个中文字符，我们需要3个字节来储存
	// utf-8建立在字符、代码点、字节三个层面上， go中的string仅仅是字节
	// 在我们的例子中，前三个字节代表着一个字符
	// 我们可以从1到4用4个字节表示一个字符(代码点, 一个代码点是32个bit), 所以用多个bit位表示任意一个字符
	// 在例子中，我用3个字节表示一个字符，因此可以读取字符按照以下方式-->
	// 3 bytes, 3 bytes, 1 bytes, 1 bytes, ... 
	str := "世界 means world"

	// UTFMax是4 --  
	var buf [utf8.UTFMax]byte

	for i, r := range str {
		r1 := utf8.RuneLen(r)

		si := i + r1

		copy(buf[:], str[i:si])

		fmt.Printf("%2d: %q; codepoint:  %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}

} 