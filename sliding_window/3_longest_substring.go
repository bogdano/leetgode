package sliding_window

func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]bool)
    length := 0
    l := 0
    for r := 0; r < len(s); r++ {
        for set[s[r]] {
            delete(set, s[l])
            l++
        }
        set[s[r]] = true
        length = max(length, r-l+1)
    }
    return length
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
