package arrays

func jump(nums []int) int {
    count := 0
    l, r := 0, 0
    for r < len(nums) - 1 {
        farthest := 0
        for i := l; i < r + 1; i++ {
            farthest = max(farthest, i + nums[i])
        }
        l = r + 1
        r = farthest
        count++
    }
    return count
}
