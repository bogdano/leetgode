package arrays

import "strings"

func reverseWords(s string) string {
    var words []string
    n := len(s)
    i := 0
    // skip leading spaces
    for i < n && s[i] == ' ' {
        i++
    }
    for i < n {
        start := i
        // move to end of word
        for i < n && s[i] != ' ' {
            i++
        }
        words = append(words, s[start:i])
        // skip spaces between words
        for i < n && s[i] == ' ' {
            i++
        }
    }
    // reverse the word slice
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}


// syntactically simpler, slower
func reverseWords_alt(s string) string {
    // trim leading and trailing spaces and reduce multiple spaces to a single space
    s = strings.TrimSpace(s)
    words := strings.Fields(s)
    // reverse the words in-place
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}
