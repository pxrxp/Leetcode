object Solution {
    def climbStairs(n: Int): Int = {

        // Fibonacci
        def noOfWays(step: Int): Int = step match {
            case 1 | 2 => step
            case _     => noOfWays(step-1) + noOfWays(step-2)
        }

        noOfWays(n)

    }
}
