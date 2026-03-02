object Solution {
    def isHappy(n: Int): Boolean = {
        
        @annotation.tailrec
        def checkHappy(slow: Int, fast: Int): Boolean =
            if (slow != fast) checkHappy(next(slow), next(next(fast)))
            else slow == 1

        @annotation.tailrec
        def next(num: Int, acc: Int = 0): Int =
            if (num <= 9) num * num + acc
            else {
                val d = num % 10
                next(num / 10, d * d + acc)
            }

        checkHappy(next(n), next(next(n)))
    }
}
