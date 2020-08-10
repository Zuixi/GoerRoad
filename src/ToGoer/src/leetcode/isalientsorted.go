package leetcode

func IsAlienSorted(words []string, order string) bool {

    // 将order定义为整形大小
    var index [26]int

    for i, val := range(order) {
        index[val - 'a'] = i
    }

    // 顺序是可以传递的， 两个两个比较就行
    for i := 0; i < len(words) - 1; i++ {
        pre := words[i]
        cur := words[i + 1]
        flag := false
        // 用第一个不同的字符进行比较
        for j := 0; j < min(len(pre), len(cur)); j++ {
            if pre[j] != cur[j] {
                if index[pre[j] - 'a'] > index[cur[j] - 'a'] {
                    return false
				}
				// 使用一个flag 证明"hel" 和"je"已经是有比较结果了，不会转到下面的if语句里
				flag = true
				break
            }
        }

		// 出现的字母都是相同的，但是长度不一样
        if len(pre) > len(cur) && !flag{
            return false
        }
    }

    return true
}
