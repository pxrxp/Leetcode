func maxDepth(s string) int {
    maximum := 0
    current := 0
    for _, c := range s {
        if c == '(' {
            current++
            if current > maximum {
                maximum = current
            }
        }
        if c == ')' {
            current--
        }
    }
    return maximum
}
