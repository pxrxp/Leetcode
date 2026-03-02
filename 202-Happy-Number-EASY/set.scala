object Solution {
    def isHappy(n: Int): Boolean = {
        def checkHappy(num: Int, seen: Set[Int]): Boolean = {
            if (num == 1) true
            else if (seen contains num) false
            else checkHappy(nextNumber(num), seen + num)
        }

        def nextNumber(num: Int): Int = {
            num.toString.map(c => c.asDigit * c.asDigit).sum
        }
    
        checkHappy(n, Set())
    }
}
