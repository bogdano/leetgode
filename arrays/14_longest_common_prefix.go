package arrays

import "strings"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    // Pre-convert all strings to []rune once
    runeStrs := make([][]rune, len(strs))
    minLength := len([]rune(strs[0]))
    for i, str := range strs {
        runes := []rune(str)
        runeStrs[i] = runes
        if len(runes) < minLength {
            minLength = len(runes)
        }
    }
    var s strings.Builder
    for i := 0; i < minLength; i++ {
        char := runeStrs[0][i]
        for j := 1; j < len(runeStrs); j++ {
            if runeStrs[j][i] != char {
                return s.String()
            }
        }
        s.WriteRune(char)
    }
    return s.String()
}
