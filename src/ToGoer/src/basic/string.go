package basic

import (
	"fmt"
	"unsafe"
	"strconv"
	"time"
	"math/rand"
	_"reflect"
)

func TestInt() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		ua uint
		ub uint8
		uc uint16
		ud uint32
		ue uint64
	)

	fmt.Println("A = ", a, "size of int = ", unsafe.Sizeof(a))
	fmt.Println("B = ", b, "size of int8 = ", unsafe.Sizeof(b))
	fmt.Println("C = ", c, "size of int16 = ", unsafe.Sizeof(c))
	fmt.Println("D = ", d, "size of int32 = ", unsafe.Sizeof(d))
	fmt.Println("E = ", e, "size of int64 = ", unsafe.Sizeof(e))

	fmt.Println("UA = ", ua, "size of int = ", unsafe.Sizeof(ua))
	fmt.Println("UB = ", ub, "size of int8 = ", unsafe.Sizeof(ub))
	fmt.Println("UC = ", uc, "size of int16 = ", unsafe.Sizeof(uc))
	fmt.Println("UD = ", ud, "size of int32 = ", unsafe.Sizeof(ud))
	fmt.Println("UE = ", ue, "size of int64 = ", unsafe.Sizeof(ue))
}

const (
	PI = 3.1515926
	E = 2.0
)

const (
	// itoa 自动从0，1，2开始递增
	Monday = 1 * iota
	Friday 
	Sunday
)

func TestConst() {
	fmt.Println("Monday = ", Monday, "\tFriday = ", Friday, "\tSunday = ", Sunday)

	var (
		n int
		f float32
		fl float64
	)

	n = 100
	f = float32(n)
	fmt.Println(n, "transfer int to float32 --> ", f)

	fl = float64(n / 3)
	fmt.Println("n / 3 transfer int to float32 --> ", fl)

	fmt.Println("tranfer float32 to int -->", int(fl * 2))

	// 溢出测试
	var (
		i64 int64
		i8 int8
	)

	i64 = 431413710704895
	i8 = int8(i64)
	fmt.Println("i64 = ", i64, "\ti8 = ", i8, "\ti64 % 255 = ", i64 % 255)

}

func TranferIntToString() {

	var (
		s string
		num = 89
	)

	// int to string
	s = strconv.Itoa(num)
	//fmt.Println("Before tranfer, type of num is ", reflect.Typeof(num))
	fmt.Println(num, " to a string is ", s)
	//fmt.Println("After tranfer, type of s is ", reflect.Typeof(s))
}

func GenerateRandString() (str string) {
	var x int64
	var res string

	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	res += strconv.FormatInt(x, 30)
	return res
}

func OperateString() {

	// 拼接字符串
	str1 := "Hello" + ",World"
	fmt.Println(str1)

	// byte = 8bits = int8
	var str2 byte
	str2 = str1[0]
	fmt.Println("str2 is ", str2)
	fmt.Printf("number of letter is %d\tletter is %c\n",str2, str2)

	// get index from 0 to 3 in str1
	str3 := str1[0:4]
	fmt.Println("str3 is ", str3)

	str4 := str1[0:len(str1)]
	fmt.Println("str4 is ", str4)

	// regard unt8 as byte
	var ub uint8
	count := 0
	for ub = 0; ub < 123; ub++ {
		count++
		fmt.Printf("(%d %c)\t", ub, ub)
		if count % 10 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")

	array := []byte(str1)
	fmt.Println("array is ", array)
	array[0] = 72
	str1 = string(array)
	fmt.Println("str1 is ", str1)
}
