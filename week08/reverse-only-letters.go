package week08

/*
题目: 仅仅反转字母
使用双指针来进行字母位置的交换
*/


func reverseOnlyLetters(s string) string {
    var charArr []byte = []byte(s)
    var length int = len(s)
    var first int = 0
    var last int = length - 1

    for last > first {
        if charArr[first] < 'A' || (charArr[first] > 'Z' && charArr[first] < 'a') || charArr[first] > 'z' {
            first++
        }
        if charArr[last] < 'A' || (charArr[last] > 'Z' && charArr[last] < 'a') || charArr[last] > 'z' {
            last--
        }
        if ((charArr[first] >= 'A' && charArr[first] <= 'Z') || (charArr[first] >= 'a' && charArr[first] <= 'z'))  && ((charArr[last] >= 'A' && charArr[last] <= 'Z') || (charArr[last] >= 'a' && charArr[last] <= 'z')) {
            var temp byte = charArr[first]
            charArr[first] = charArr[last]
            charArr[last] = temp
            first++
            last--
        }
    }
    return string(charArr)
}

