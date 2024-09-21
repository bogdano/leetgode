package arrays

import "strings"

func lengthOfLastWord(s string) int {
    var str strings.Builder
    prevSpace := false
    for _, char := range(s) {
        if char != ' ' {
            if prevSpace == true {
                str.Reset()
                prevSpace = false
            }
            str.WriteRune(char)
        } else if char == ' ' {
            prevSpace = true
        }
    }
    return str.Len()
}
