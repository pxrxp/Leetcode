object Solution {
    def plusOne(digits: Array[Int]): Array[Int] = {
        
        @annotation.tailrec
        def inc(lastIndex: Int) : Array[Int] = {
            if (lastIndex == -1) 1 +: digits
            else {
                if (digits(lastIndex) == 9) {
                    digits(lastIndex) = 0
                    inc(lastIndex - 1)
                } else {
                    digits(lastIndex) += 1
                    digits
                }
            }
        }

        inc(digits.length-1)
    }
}
