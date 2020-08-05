package leetcode

import (
	"strconv"
	"strings"
)

// #788
func RotatedDigits(N int) int {
    // 0 1 8 -> 0 1 8
    // 2 5 -> 5 2 
    // 6 9 -> 9 6
    // X != R(X)
    // solution 
    count := 0
    for  i := 1; i <= N; i++ {
        if IsGoodNum(i) {
            count++
        }
    }
    return count
}

func IsGoodNum(n int) bool {
    s := strconv.FormatInt(int64(n), 10)

    if strings.Contains(s, "3") || strings.Contains(s, "4") || strings.Contains(s, "7") {
        return false
    }
	
	if !strings.Contains(s, "2") && !strings.Contains(s, "5") && !strings.Contains(s, "6") && !strings.Contains(s, "9"){
        return false
    }

    return true
}