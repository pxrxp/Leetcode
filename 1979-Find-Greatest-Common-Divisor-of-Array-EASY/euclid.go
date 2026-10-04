func findGCD(nums []int) int {
    max := 0
    min := 1001

    for _, n := range nums {
        if n > max {
            max = n
        }
        if n < min {
            min = n
        }
    }

    return gcd(min, max)
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}
