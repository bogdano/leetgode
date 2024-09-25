package two_pointers

import "sort"

func threeSum(nums []int) [][]int {
    var combos [][]int
    sort.Ints(nums)
    for i, num := range(nums) {
        // skip identical values, to prevent duplicates in set
        if i > 0 && num == nums[i-1] {
            continue
        }
        // set next two pointers
        l, r := i+1, len(nums)-1
        for l < r {
            sum := num + nums[l] + nums[r]
            if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            } else {
                combos = append(combos, []int{num, nums[l], nums[r]})
                // update one of the pointers until a new number is reached
                l++
                for nums[l] == nums[l-1] && l < r {
                    l++
                }
            }
        }
    }
    return combos
}
