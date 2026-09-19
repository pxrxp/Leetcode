func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for _, v := range nums[1:] {
		currentSum = max(currentSum+v, v)
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
