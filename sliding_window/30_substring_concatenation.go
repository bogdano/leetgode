package sliding_window

func findSubstring(s string, words []string) []int {
    if len(words) == 0 || len(s) == 0 {
        return []int{}
    }
    wordLength := len(words[0])
    totalWords := len(words)
    // make a map with wordcounts of the given words slice
    wordsCount := make(map[string]int)
    for _, word := range words {
        wordsCount[word] += 1
    }
    var indices []int
    // loop over starting positions up to wordLength
    for offset := 0; offset < wordLength; offset++ {
        left, right := offset, offset
        currentCount := make(map[string]int)
        count := 0

        for right+wordLength <= len(s) {
            word := s[right:right+wordLength]
            right += wordLength
            if _, found := wordsCount[word]; found {
                currentCount[word] += 1
                count++
                for currentCount[word] > wordsCount[word] {
                    leftWord := s[left:left+wordLength]
                    currentCount[leftWord] -= 1
                    count--
                    left += wordLength
                }
                // if the count is good, we found a substring: record the index
                if count == totalWords {
                    indices = append(indices, left)
                }
            } else {
                // reset the counters
                currentCount = make(map[string]int)
                count = 0
                left = right
            }
        }
    }
    return indices
}
