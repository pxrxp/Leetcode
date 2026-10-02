func networkDelayTime(times [][]int, n int, k int) int {
    cost := make([]int, n)
    for i := 0; i < n; i++ {
        cost[i] = 1e3
    }
    cost[k-1] = 0

    for i := 0; i < n-1; i++ {
        changed := false
        for _, edge := range times {
            source, target, time := edge[0], edge[1], edge[2]
            if cost[target-1] > cost[source-1] + time {
                cost[target-1] = cost[source-1] + time
                changed = true
            }
        }
        if !changed {
            break
        }
    }

    max := cost[0]
    for _, c := range cost {
        if c == 1e3 {
            return -1
        }
        
        if max < c {
            max = c
        }
    }
    return max
}
