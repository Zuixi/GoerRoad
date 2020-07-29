package leetcode

func FindTheDifference(s string, t string) byte {
    
    // 插入字母有可能不是第一次; byte 对标 char; string的类型是rune
    msg := make(map[byte]int, len(t) + 1)

    for _, ch := range(s) {
        msg[byte(ch)]++
    }

    for _, ch := range(t) {
        msg[byte(ch)]--
    }

    for ch, val := range(msg) {
        if val == -1 {
            return ch
        }
    }
    return 'a'
}