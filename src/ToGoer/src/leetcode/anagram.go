package leetcode

// 字母异位词
// 两个单词如果包含同样的字母，并且次序不同，那么这两个单词就是字母异位词
// slient 和 listen 是异位词， 而apple 和aplee不是
func IsAnagram(s string, t string) bool {
	ls := len(s)
    lt := len(t)

    if ls != lt {
        return false
    }

    // 处理两个特殊case 两个都是空字符串或者都是一个字母的时候
    if ls == lt {
        if (ls == 0) {
            return true
        }
        if ls == 1 && s == t {
            return true
        }
    }

    // 使用map 对构成单词的字符进行处理
    ms := make(map[rune]int)
    for _, val := range(s) {
        ms[val]++
    }

    for _, val := range(t) {
        ms[val]--
    }

    flag := true
    for _, count := range(ms) {
        if count != 0 {
            flag = false
            break
        }
    }

    if flag && s != t {
        return true
    } else {
        return false
    }
}


func IsAnagramPlus(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 使用一个数组来充当map的使用，方法很好
	count := [26]int{}

	for _, val := range(s) {
		count[val - 'a']++
	}

	for _, val := range(t) {
		count[val - 'a']--
		if (count[val - 'a'] < 0) {
			return false
		}
	}
	return true
}