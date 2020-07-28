package basic

import (
	"strings"
	"fmt"
)

// bitwise AND &
// &可以用来判断一个数是否是奇数还是偶数，和1做AND运算，结果为1就是奇数，0是偶数
func And(a, b uint16) uint16 {
	return a & b
}

func JudgeNumber(a uint16) bool {
	if a & 1 == 1 {
		return false
	}
	return true
}

// OR |
// OR(a, b) = 1; when a = 1 or b = 1 else = 0
// 可以利用OR为给定的整数有选择地设置单个位,比如讲一个数的某些位置0，其余位置1

const (
	UPPER = 1 // 大写字符串
	LOWER = 2 // 小写字符串
	CAP   = 4 //首字母大写
	REV   = 8 // 反转字符串
)

func procstr(str string, conf byte) string {
	// 反转字符 匿名函数
	rev := func (s string ) string {
		runes :=  []rune(s)
		n := len(runes)

		for i := 0; i < n / 2; i++ {
			// 直接交换两个值
			runes[i], runes[n - 1 - i] = runes[n - 1 - i], runes[i]
		}

		return string(runes)
	}

	//查询配置中的位操作
	if conf & UPPER != 0 {
		str = strings.ToUpper(str)
	}
	if conf & LOWER != 0 {
		str = strings.ToLower(str)
	}
	if conf & CAP != 0 {
		str = strings.Title(str)
	}
	if conf & REV != 0 {
		str = rev(str)
	}

	return str
}

func OrString() string {
	return procstr("HELLO PEOPLE", LOWER | REV | CAP)
	// should return string as follows
	// hello people firstly
	// elpoep olleh secondly
	// Elpoep olleh finally
}

// XOR ^ 
// XOR(a, b) = 1; only if a != b else = 0
// XOR 可以把二进制位的一个值变成另外一个值，例如 0xCEFF ^ 0xFF00
// XOR 还可以用来判断两个数字的符号是否一样，若 a ^ b >= 0 或者a ^ b < 0为true的时候，a和b具有相同的符号

func IsSameSignal(a, b int) bool {
	if a ^ b >= 0 {
		return true
	} else {
		return false
	}
}

// ^ 也可以作为取反运算符
func ReverseInt(a int) int {
	return ^a
}

// &^ 操作符 与非 
// Given operands a, b
// AND_NOT(a, b) == AND(a, NOT(b))
// AND_NOT(a, 1) = 0; clears a
// AND_NOT(a, 0) = a
// 如果第二个操作数是1，那么可以清除第一个操作数的位

func TestAndNot() {
	var a byte = 0xAB;
	fmt.Print("%08b\n", a)

	a &^= 0x0F
	fmt.Print("%08b\n", a)

	// 将1010 1011 变成1010 0000
	var b byte = 0xAB
	b &^= 0x0F
	fmt.Print("%08b\n", b)

}

// 移位运算符 << and >> 
// a << n --> shifts all bits in a to the left n times
// a >> n --> shifts all bits in a to the right n times

func TestRL() {
	// 将0x00000011左移三次
	var a int8 = 3
	fmt.Println("%08b\n", a)
	fmt.Println("%08x\n", a << 1)
	fmt.Println("%08x\n", a << 2)
	fmt.Println("%08x\n", a << 3)
}