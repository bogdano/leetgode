package arrays

import "math/rand"

type RandomizedSet struct {
    nums       []int
    indices    map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet {
        nums:   []int{},
        indices: make(map[int]int),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.indices[val]; exists {
        return false
    }
    this.nums = append(this.nums, val)
    this.indices[val] = len(this.nums) - 1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    index, exists := this.indices[val]
    if !exists {
        return false
    }
    lastIndex := len(this.nums) - 1
    lastVal := this.nums[lastIndex]
    this.nums[index] = lastVal
    this.indices[lastVal] = index
    this.nums = this.nums[:lastIndex]
    delete(this.indices, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    index := rand.Intn(len(this.nums))
    return this.nums[index]
}
