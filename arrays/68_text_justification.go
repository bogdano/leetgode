package arrays

import "strings"

func fullJustify(words []string, maxWidth int) []string {
    var result []string
    var line []string
    var sb strings.Builder
    length := 0
    i := 0
    for i < len(words) {
        // if length of the words in the line plus the spaces between them plus the next word is too big
        if (length + len(line) + len(words[i])) > maxWidth {
            extra_spaces := maxWidth - length
            // number of space between each word
            spaces := extra_spaces / max(1, len(line)-1)
            // the remainder spaces to be allotted greedily to left words
            remainder := extra_spaces % max(1, len(line)-1)
            // iterate over words in current line to add spaces
            for j := 0; j < max(1, len(line)-1); j++ {
                // write the current word
                sb.WriteString(line[j])
                // add necesary spaces after each word
                for range(spaces) {
                    sb.WriteRune(' ')
                }
                // if remainder space, greediliy add after first words
                if remainder != 0 {
                    sb.WriteRune(' ')
                    remainder--
                }
                // return the full line to the lines slice
                line[j] = sb.String()
                sb.Reset()
            }
            // build the line as a string
            for _, word := range(line) {
                sb.WriteString(word)
            }
            result = append(result, sb.String())
            sb.Reset()
            // reset the line slice and length for next line processing
            line, length = nil, 0
        }
        // append the overflow word to the next line
        line = append(line, words[i])
        length += len(words[i])
        i++
    }
    // now populate the last line
    for i, word := range(line) {
        sb.WriteString(word)
        // don't add space after last word
        if i < len(line)-1 {
            sb.WriteRune(' ')
        }
    }
    // add spaces to fill last line
    trail_space := maxWidth - sb.Len()
    for range(trail_space) {
        sb.WriteRune(' ')
    }
    result = append(result, sb.String())
    sb.Reset()
    return result
}
