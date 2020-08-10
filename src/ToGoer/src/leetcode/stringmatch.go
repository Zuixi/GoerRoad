package leetcode

import (
	"strings"
)

func StringMatching(words []string) []string {
    cap := len(words)
    count := 0
    res := make([]string, cap)
    
    var flag [101]bool

    for i := 0; i < cap; i++ {
        for j := 0; j < cap; j++ {
            if !flag[i] && words[i] != words[j] && strings.Contains(words[j],words[i]) {
                res[count] = words[i]
                flag[i] = true
                count++
            }
        }
    }
    res = res[0 : count]
    return res
}