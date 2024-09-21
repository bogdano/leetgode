package arrays

func trap(height []int) int {
    n := len(height)
    leftMax := make([]int, n)
    currMax := 0
    // populate slice holding max values between current index and left index
    for i := 1; i < n; i++ {
        if height[i] < height[i-1] {
            currMax = max(currMax, height[i-1])
        }
        leftMax[i] = currMax
    }
    rightMax := make([]int, n)
    currMax = 0
    // populate slice holding max values between current index and right index
    for i := n-2; i >=0; i-- {
        if height[i] < height[i+1] {
            currMax = max(currMax, height[i+1])
        }
        rightMax[i] = currMax
    }
    total_water := 0
    for i := 0; i < n; i++ {
        water := min(leftMax[i], rightMax[i])
        water -= height[i]
        if water < 0 {
            water = 0
        }
        total_water += water
    }
    return total_water
}

// alternative implementation, using two pointers (less memory)
func trap_alt(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }
    l, r := 0, n-1
    leftMax, rightMax := height[l], height[r]
    result := 0
    for l < r {
        if leftMax < rightMax {
            l++
            leftMax = max(leftMax, height[l])
            result += leftMax - height[l]
        } else {
            r--
            rightMax = max(rightMax, height[r])
            result += rightMax - height[r]
        }
    }
    return result
}
