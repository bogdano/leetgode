package arrays

func removeElement(nums []int, val int) int {
    i := 0
    for _, number := range nums {
        if number != val {
            nums[i] = number
            i++
        }
    }
    return i
}
