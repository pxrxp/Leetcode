// Imagine we have a vessel of infinite capacity
// We can fill any number of x or y jugs
// We can subtract any number of x or y jugs from that vessel
// We need to check whether ax+by=target is possible
// We cant fill 0.1, 0.2 ... (decimal)
// So, it is a diophantine equation
// GCD(a,b) means the highest number that divides both a and b
// So, if g = GCD(a,b), a = a'g and b = b'g
// (a-b) = (a'-b')g is also divisible by g
// Hence, any target must be divisible by g to be valid
// But since, we dont have infinite capacity,
// we also check if target can fit if both jugs combined.

func canMeasureWater(x int, y int, target int) bool {
	return target <= x+y && target%gcd(x, y) == 0
}

func gcd(x, y int) int {
	var a, b int
	if x > y {
		a, b = y, x
	} else {
		a, b = x, y
	}

	for a > 0 {
		a, b = b%a, a
	}

	return b
}
