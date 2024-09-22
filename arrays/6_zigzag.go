package arrays

import "strings"

func convert(s string, numRows int) string {
    n := len(s)
    if numRows == 1 {
        return s
    }
    var res strings.Builder
    for i := 0; i < numRows; i++ {
    	// increment is the distance between the first character of the current row and the first character of the next row
        increment := 2 * (numRows-1)
        for j := i; j < n; j += increment {
            res.WriteByte(s[j])
            if i > 0 && i < numRows-1 && (j + increment - 2*i) < n {
                res.WriteByte(s[j + increment - 2 * i])
            }
        }
    }
    return res.String()
}
