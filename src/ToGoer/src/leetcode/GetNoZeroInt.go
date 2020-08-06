package leetcode

import (
	"strconv"
	"strings"
)

// #1317
func GetNoZeroIntegers(n int) []int {
	rs := make([]int, 2)
	count := 1
    for {
        if !strings.Contains(strconv.Itoa(count), "0") && !strings.Contains(strconv.Itoa(n - count), "0") {
            rs[0] = count
            rs[1] = n - count
            break
        }
        count++
    }
    return rs
}