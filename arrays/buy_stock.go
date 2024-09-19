package arrays

func maxProfit(prices []int) int {
    i := 0
    profit := 0
    for j := 1; j < len(prices); j++ {
        if prices[i] < prices[j] {
            profit = max(profit, prices[j] - prices[i])
        } else {
            i = j
        }
    }
    return profit
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
