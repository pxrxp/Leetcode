func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, target+1)
	for i := range dp {
		dp[i] = make([]bool, len(nums)+1)
	}
	for i := range dp { // Target = i
		for j := range dp[i] { // Use arr up to index j
			if i == 0 {
				dp[i][j] = true
			} else if j == 0 {
				dp[i][j] = false
			} else {
				choose := false
				dontChoose := dp[i][j-1]

				if i >= nums[j-1] {
					choose = dp[i-nums[j-1]][j-1]
				}

				dp[i][j] = choose || dontChoose
			}
		}
	}
	return dp[target][len(nums)]
}
