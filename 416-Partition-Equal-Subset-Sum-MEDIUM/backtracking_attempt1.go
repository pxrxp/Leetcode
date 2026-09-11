func backtrack(nums []int, i int, sum int, target int) bool {
	if sum == target {
		return true
	}
	if sum > target {
		return false
	}
	if i >= len(nums) {
		return false
	}

	if backtrack(nums, i+1, sum+nums[i], target) {
		return true
	}
	return backtrack(nums, i+1, sum, target)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	return backtrack(nums, 0, 0, sum/2)
}
