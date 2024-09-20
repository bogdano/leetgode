package arrays

func removeDuplicates_2(nums []int) int {
    i := 0
    ctr := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
            ctr = 0
        } else {
            if ctr < 1 {
                i++
                nums[i] = nums[j]
                ctr++
            }
        }
    }
    return i+1
}
