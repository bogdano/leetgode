package sliding_window

func minWindow(s string, t string) string {
    if t == "" {
        return t
    }
    t_map, window := make(map[rune]int), make(map[rune]int)
    for _, letter := range t {
        t_map[letter] += 1
    }

    have, need := 0, len(t_map)
    result, resLen := []int{-1,-1}, len(s)+1
    l := 0
    for r, letter := range s {
        window[letter] += 1
        if _, found := t_map[letter]; found && window[letter] == t_map[letter] {
            have++
        }
        for have == need {
            if r-l+1 < resLen {
                result = []int{l, r}
                resLen = r-l+1
            }
            window[rune(s[l])] -= 1
            if _, found := t_map[rune(s[l])]; found && window[rune(s[l])] < t_map[rune(s[l])] {
                have--
            }
            l++
        }
    }
    if resLen < len(s)+1 {
        return s[result[0]:result[1]+1]
    } else {
        return ""
    }
}
