package arrays

func strStr(haystack string, needle string) int {
    hay := []rune(haystack)
    pin := []rune(needle)
    h_len, p_len := len(hay), len(pin)
    if p_len == 0 {
        return 0
    }

    for i := 0; i <= h_len-p_len; i++ {
        match := true
        for j := 0; j < p_len; j++ {
            if hay[i+j] != pin[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }

    return -1
}
