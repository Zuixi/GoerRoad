package basic

import (
	"fmt"
)

func TestEnum() {
	// 在go中是没有直接支持枚举的关键字的
	// 但是我们可以通过const + iota来实现枚举

	type MoneyType int

	const (
		RMB MoneyType = iota + 1 // 以1开始自增
		DOLLAR					 // value --> 2
		JAPANESE				 // value --> 3
	)

	printMoney := func(this MoneyType) {
		switch this {
		case RMB:
			fmt.Println("you select RMB")
			break
		case DOLLAR:
			fmt.Println("you select dollar")
			break
		case JAPANESE:
			fmt.Println("you select japanese")
			break
		default:
			fmt.Println("there is no type you selected")
			break		
		}
	}

	printMoney(DOLLAR)
	fmt.Println()
	printMoney(RMB)
	fmt.Println()
	printMoney(4)
}