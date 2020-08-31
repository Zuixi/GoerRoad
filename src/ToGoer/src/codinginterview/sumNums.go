package codinginterview

// ----------------------------------------------------------------------------------------------
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
// ----------------------------------------------------------------------------------------------
func sumNumsBase(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return n + sumNumsBase(n - 1)
}

// 不适用for循环的话，则可以使用递归代替
// 递归的出口不能使用if 语句
// 使用别的方法来确定出口，那就是逻辑短路 &&
func SumNumsPlus(n int) int {
	res := 0
	
	// 使用匿名函数做递归
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n - 1)
	}
	sum(n)
	return res
}


