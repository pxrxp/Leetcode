func nextPermutation(nums []int)  {
    l := len(nums)-1
    if l-1 < 0 {
        return
    }

    i := l-1    // Replace
    j := l      // Replace with
    for ; i >= 0 && nums[i] >= nums[i+1]; i-- {}
    if i >= 0 {
        for ; j >= 0 && nums[j] <= nums[i]; j-- {}
        nums[i], nums[j] = nums[j], nums[i]
    }
    for k, kk := i+1, l; k < kk; k, kk = k+1, kk-1 {
        nums[k], nums[kk] = nums[kk], nums[k]
    }
}
