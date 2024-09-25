package two_pointers

func maxArea(height []int) int {
    l, r := 0, len(height)-1
    water := 0
    for l < r {
        if height[r] > height[l] {
            water = max(water, height[l]*(r-l))
            l++
        } else {
            water = max(water, height[r]*(r-l))
            r--
        }
    }
    return water
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
