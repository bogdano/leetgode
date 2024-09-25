package two_pointers

func isSubsequence(s string, t string) bool {
    i := 0
    if len(s) == 0 {
        return true
    }
    for j := 0; j < len(t); j++ {
        if i == len(s)-1 && s[i] == t[j] {
            return true
        }
        if s[i] == t[j] {
            i++
        }
    }
    return false
}
