package arrays

func maxProfit_2(prices []int) int {
    profit := 0
    for j := 1; j < len(prices); j++ {
        if prices[j] > prices[j-1] {
            profit += prices[j]-prices[j-1]
        }
    }
    return profit
}
