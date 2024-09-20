package arrays

func candy(ratings []int) int {
    total_candy := 0
    n := len(ratings)
    candies := make([]int, n)
    for i := 0; i < n; i++ {
        candies[i] = 1
    }
    // now go left to right
    for i := 1; i < n; i++ {
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1] + 1
        }
    }
    //now go right to left
    for i := n-2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            candies[i] = max(candies[i], candies[i+1] + 1)
        }
    }
    // now sum up the candies
    for i := 0; i < n; i++ {
        total_candy += candies[i]
    }
    return total_candy
}
