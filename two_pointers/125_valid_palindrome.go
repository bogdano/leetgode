package two_pointers

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
    s = preprocessString(s)
    n := len(s)
    for i := 0; i < n / 2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func preprocessString(s string) string {
    var sb strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            sb.WriteRune(unicode.ToLower(r))
        }
    }
    return sb.String()
}
