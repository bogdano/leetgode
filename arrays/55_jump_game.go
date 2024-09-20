package arrays

func canJump(nums []int) bool {
    maxReach := 0
    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            return false
        }
        maxReach = max(maxReach, i+nums[i])
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    return true
}

// alternative, slightly faster
func canJump_alt(nums []int) bool {
    goalPost := len(nums)-1
    for i := len(nums)-1; i >= 0; i-- {
        if i + nums[i] >= goalPost {
            goalPost = i
        }
    }
    if goalPost == 0 {
        return true
    } else {
        return false
    }
}
