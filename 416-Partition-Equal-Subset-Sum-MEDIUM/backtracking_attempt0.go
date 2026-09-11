func sum(indices map[int]struct{}, arr []int) int {
	sum := 0
	for val := range indices {
		sum += arr[val]
	}
	return sum
}

func check(indices map[int]struct{}, arr []int, totalSum int) bool {
	return (sum(indices, arr)*2 == totalSum)
}

func backtrack(indices map[int]struct{}, arr []int, totalSum int) bool {
	if check(indices, arr, totalSum) {
		return true
	}

	for i := range arr {
		_, ok := indices[i]
		if !ok {
			indices[i] = struct{}{}
			if backtrack(indices, arr, totalSum) {
				return true
			}
			delete(indices, i)
		}
	}
	return false
}

func canPartition(nums []int) bool {
	totalSum := 0
	for _, val := range nums {
		totalSum += val
	}
	indices := make(map[int]struct{})
	return backtrack(indices, nums, totalSum)
}
