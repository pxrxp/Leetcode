func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)

    for i := range dp {
        dp[i] = make([]int, len(word2)+1)
        for j := range dp[i] {
            if i == 0 {
                dp[i][j] = j
            } else if j == 0 {
                dp[i][j] = i
            } else if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = 1 + min(
                    dp[i-1][j],    // Delete a char
                    dp[i][j-1],    // Insert a char
                    dp[i-1][j-1],  // Replace a char
                )
            }
        }
    }

    return dp[len(word1)][len(word2)]
}

func min(nums ...int) int {
    m := nums[0]
    for _, v := range nums {
        if v < m {
            m = v
        }
    }
    return m
}
