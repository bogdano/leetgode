package matrix

func spiralOrder(matrix [][]int) []int {
    width, height := len(matrix[0]), len(matrix)
    var result []int
    // right, left, top, bottom pointers
    l, r := 0, width
    t, b := 0, height
    for l < r && t < b {
        // scoop along top
        for i := l; i < r; i++ {
            result = append(result, matrix[t][i])
        }
        t++
        // scoop down far right col
        for i := t; i < b; i++ {
            result = append(result, matrix[i][r-1])
        }
        r--
        // if pointers overlap, end it
        if !(l < r && t < b) {
            break
        }
        // scoop along bottom
        for i := r-1; i >= l; i-- {
            result = append(result, matrix[b-1][i])
        }
        b--
        // scoop up far left col
        for i := b-1; i >= t; i-- {
            result = append(result, matrix[i][l])
        }
        l++
    }
    return result
}
