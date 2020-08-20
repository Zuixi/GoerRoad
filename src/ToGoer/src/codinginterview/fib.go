package codinginterview

// #leetcode codinginterview 7
func Fib(n int) int {
    var f0 = 0
    var f1 = 1
	var i = 2
	
    for {
        if i > n {
            break
        }

        f1 = f0 + f1
        f0 = f1 - f0
        // 这里需要注意题意，超过某个数的时候需要重新取余
        if f1 > (1e9 + 7) {
            f1 = f1 % (1e9 + 7)
        }
        
        if f0 > (1e9 + 7) {
            f0 = f0 % (1e9 + 7)
        }
        i++
    }

    if n == 0 {
		return f0
	} else {
		return f1
	}
}

// 递归解法，但是求解速度慢，效率高，可以使用剪枝法优化
// 在实际提交代码时，这种解法会超时
func RecurbFib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return RecurbFib(n - 1) + RecurbFib(n - 2)
}

// 动态规划题解
func DynamicFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i - 1] + dp[i - 2]) % (1e9 + 7)
	}

	return dp[n]
}