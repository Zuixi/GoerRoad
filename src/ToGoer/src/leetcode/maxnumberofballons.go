package leetcode

// #1189
func MaxNumberOfBalloons(text string) int {
    var rs [26]int

    // 记录字母出现的次数
    for _, val := range text {
        rs[val - 'a']++
    }

    count := 0
    for {
        if rs['b' - 'a'] >= 1 && rs['a' - 'a'] >= 1 && rs['n' - 'a'] >= 1 {
            if rs['l' - 'a'] >= 2 && rs['o' - 'a'] >= 2 {
                count++

                rs['b' - 'a']--
                rs['a' - 'a']--
                rs['n' - 'a']--
                rs['l' - 'a'] -= 2
                rs['o' - 'a'] -= 2
            } else {
                break
            }
        } else {
            break
        }
    }

    return count
}

func  MaxNumberOfBalloonsPro(text string) int {
    var rs [26]int

    // 记录字母出现的次数
    for _, val := range text {
        rs[val - 'a']++
    }

    count := 0
    
    count = min(len(text),rs[0])
    count = min(count,rs[1])
    count = min(count,rs['n' - 'a'])
    count = min(count,rs['l' - 'a'] / 2)
    count = min(count,rs['o' - 'a'] / 2)
    
    return count
}

func min(x , y int) int {
    if x > y {
        return y
    }
    return x
}