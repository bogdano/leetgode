package sliding_window

func minSubArrayLen(target int, nums []int) int {
    minLen := len(nums)+1 // set to invalid subarray length
    sum, left := 0, 0
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            if minLen > right - left + 1 {
                minLen = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }
    if minLen == len(nums) + 1 {
        return 0
    }
    return minLen
}
