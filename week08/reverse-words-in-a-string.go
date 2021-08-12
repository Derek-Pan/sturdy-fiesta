package week08

/*

题目: 翻转字符串里的单词
1. 先讲原字符串切割成多个单词
2. 单词反转拼接 
*/

func reverseWords(s string) string {
    var words []string = make([]string, 0, 0)
    var charArr []byte = make([]byte, 0, 0)
    var ans string

    for i := 0; i < len(s); i++ {
        if s[i] == ' ' && len(charArr) > 0 {
            words = append(words, string(charArr))
            charArr = charArr[:0]
        } else if s[i] != ' ' {
            charArr = append(charArr, s[i])
        }
    }
    if len(charArr) > 0 {
        words = append(words, string(charArr))
    }
    var length int = len(words)
    ans = words[length - 1]
    for i := length - 2; i >= 0; i-- {
        ans = ans + " " + words[i]
    }
    return ans
}
