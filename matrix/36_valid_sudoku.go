package matrix

func isValidSudoku(board [][]byte) bool {
    rows := make(map[int]map[byte]bool)
    cols := make(map[int]map[byte]bool)
    boxes := make(map[int]map[byte]bool)
    for i := range 9 {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }
    for i := range 9 {
        for j := range 9 {
            num := board[i][j]
            if num == '.' {
                continue
            }
            box := (i/3)*3 + j/3
            if rows[i][num] || cols[j][num] || boxes[box][num] {
                return false
            }
            rows[i][num] = true
            cols[j][num] = true
            boxes[box][num] = true
        }
    }
    return true
}
