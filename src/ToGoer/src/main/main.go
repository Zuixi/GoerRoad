package main

import (
	"fmt"
	"ToGoer/src/basic" 
)

func main() {
	fmt.Println("hello")
	basic.TestInt()

	fmt.Println()
	basic.TranferIntToString()

	fmt.Println("Generate rand string as follows")
	res := basic.GenerateRandString()
	fmt.Println("return string is ", res)

	basic.OperateString()
}