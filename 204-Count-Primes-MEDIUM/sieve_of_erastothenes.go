func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }

    primes := make([]bool, n)
    for i:=2; i<n; i++ {
        primes[i] = true
    }
    count := n-2

    for p:=2; p*p<=n; p++ {
        if primes[p] {
            for i:=p*p; i<n; i+=p {
                if primes[i] {
                    primes[i] = false
                    count--
                }
            }
        }
    }

    return count
}
