func recurse(coins []int, amount int, dp map[int]int) int {
	if v, ok := dp[amount]; ok {
		return v
	}

	min := math.MaxInt32
	for _, coin := range coins {
		if amount >= coin {
			noOfCoinsNeeded := recurse(coins, amount-coin, dp)
			if noOfCoinsNeeded != -1 && noOfCoinsNeeded+1 < min {
				min = noOfCoinsNeeded + 1
			}
		}
	}

	if min != math.MaxInt32 {
		dp[amount] = min
	} else {
		dp[amount] = -1
	}

	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	dp := map[int]int{0: 0}
	ans := recurse(coins, amount, dp)
	return ans
}
