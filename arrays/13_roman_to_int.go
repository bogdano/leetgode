package arrays

func romanToInt(s string) int {
    numerals := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    n := len(s)
    var result int
    for i := 0; i < n; i++ {
        if i+1 < n && numerals[s[i]] < numerals[s[i+1]] {
            result -= numerals[s[i]]
        } else {
            result += numerals[s[i]]
        }
    }
    return result
}
