func findNumbers(nums []int) int {
    count := 0
    for _, v := range nums {
        d := 0
        for v > 0 {
            d += 1
            v /= 10
        }
        if d % 2 == 0 {
            count += 1
        }
    }
    return count
}
