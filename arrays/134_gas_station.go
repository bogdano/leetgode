package arrays

func canCompleteCircuit(gas []int, cost []int) int {
    gas_tank := 0
    start := 0
    if sum(gas) < sum(cost) {
        return -1
    }
    for i := 0; i < len(gas); i++ {
        gas_tank += (gas[i] - cost[i])
        if gas_tank < 0 {
            gas_tank = 0
            start = i + 1
        }
    }
    return start
}

func sum(nums []int) int {
    sum := 0
    for _, num := range(nums) {
        sum += num
    }
    return sum
}
