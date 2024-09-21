package arrays

import "strings"
func intToRoman(num int) string {
    numerals := []struct {
        Value  int
        Symbol string
    }{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }
    var result strings.Builder
    for _, numeral := range(numerals) {
        for num >= numeral.Value {
            result.WriteString(numeral.Symbol)
            num -= numeral.Value
        }
    }
    return result.String()
}

// using slice of structs here because iterating over a map has unforeseeable ordering
