func backtrack(nums []int, i int, ans [][]int, j int) [][]int {
	if i >= len(nums) {
		return ans
	}

	choose := slices.Clone(ans[j])
	choose = append(choose, nums[i])

	ans = append(ans, choose)
	newJ := len(ans) - 1

	ans = backtrack(nums, i+1, ans, newJ)
	ans = backtrack(nums, i+1, ans, j)

	return ans
}

func subsets(nums []int) [][]int {
	l := float64(len(nums))
	ans := make([][]int, 0, int(math.Pow(2, l)))
	ans = append(ans, []int{})
	ans = backtrack(nums, 0, ans, 0)
	return ans
}
