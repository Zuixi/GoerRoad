package leetcode

import (
	"strconv"
)

// 类似于 strconv.FormatInt(n int64, base int) string
func GetStringOfInt(n int, base int) string {
	if n < 0 {
		return ""
	}

	res := ""
	for {
		// transform num to string
		// transform string to num with strconv.Atoi()
		s := strconv.Itoa(n % base)
		res = res + s
		n = n / base
		if n == 0 {
			break
		}
	}

	return res
}

// 861
func BinaryGap(N int) int {
    res := strconv.FormatInt(int64(N), 2)
    distance := 0
    for i := 0; i < len(res); i++ {
        if res[i] == '1' {
            for j := i + 1; j < len(res); j++ {
                if res[j] == '1'{
                    m := j - i
                    if m > distance {
                        distance = m
					}
					break
                }                  
            }
        }        
    }
    return distance
}