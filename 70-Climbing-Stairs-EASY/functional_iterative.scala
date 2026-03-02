object Solution {
    def climbStairs(n: Int): Int = {

        def noOfWays(step: Int, prev: Int = 0, prevPrev: Int = 0): Int = step match {
            case 1 => step
            case _ =>
                (2 until n).foldLeft((1,2)) {
                    case ((a,b), _) => (b, a+b)
                }._2
        }

        noOfWays(n)

    }
}
