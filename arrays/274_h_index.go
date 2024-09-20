package arrays

func hIndex(citations []int) int {
    var papers int
    h := len(citations)
    paper_counts := make([]int, h+1)
    for _, c := range(citations) {
        paper_counts[min(h, c)] += 1
    }
    papers = paper_counts[h]
    for papers < h {
        h--
        papers += paper_counts[h]
    }
    return h
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
