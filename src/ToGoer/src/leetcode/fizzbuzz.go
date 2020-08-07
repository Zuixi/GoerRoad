package leetcode

import (
	"strconv"
)
const (
    Thr string = "Fizz"
    Five string = "Buzz"
    Both string = "FizzBuzz"
)

// #412
func FizzBuzz(n int) []string {
    res := make([]string, n)

    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            res[i - 1] = Both
        } else if i % 5 == 0 {
            res[i - 1] = Five
        } else if i % 3 == 0 {
            res[i - 1] = Thr
        } else {
            res[i - 1] = strconv.Itoa(i)
        }
    }

    return res
}