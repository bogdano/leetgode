package arrays

func rotate(nums []int, k int)  {
    k = k % len(nums)
    count := 0
    for start := 0; count < len(nums); start++  {
        current := start
        prev := nums[start]
        for {
            next := (current + k) % len(nums)
            nums[next], prev = prev, nums[next]
            current = next
            count++
            if start == current {
                break
            }
        }
    }
}

// different solution, faster but more memory

func rotate_alt(nums []int, k int)  {
    n := len(nums)
    k = k % n
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
