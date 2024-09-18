package arrays

func majorityElement(nums []int) int {
    var m int
    c := 0
    for _, number := range(nums) {
        if c == 0 {
            m = number
            c = 1
        } else if m == number {
            c++
        } else {
            c--
        }
    }
    return m
} // boyer-moore
