func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    cost := make([]int, n)
    for i := range cost {
        cost[i] = 1e5
    }
    cost[src] = 0

    // Each iteration adds a stop
    for i := 0; i < k+1; i++ {
        // Copy is so that we dont modify cost while iteration going on
        nextCost := make([]int, n)
        copy(nextCost, cost)

        for _, flight := range flights {
            from, to, price := flight[0], flight[1], flight[2]
            if nextCost[to] > cost[from] + price {
                nextCost[to] = cost[from] + price
            }
        }

        copy(cost, nextCost)
    }

    if cost[dst] == 1e5 {
        return -1
    } else {
        return cost[dst]
    }
}
