package codinginterview

// f(x) 表示到x级台阶的方法数， 最后一步只能是一步或者两步，所以可以列出如下表达式
// f(x) = f(x - 1) + f(x - 2)
// 当台阶只有一级的时候f(1) = 1, 显然f(0) = 0
// 所以我们可以用上面的递推式求出n级台阶的方法数

func NumWays(n int) int {
	pre, cur, res := 0, 0, 1

	for i := 1; i <= n; i++ {
		pre = cur
		cur = res
		res = pre + cur
	}

	return res
}