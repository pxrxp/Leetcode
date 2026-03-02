object Solution {
    def plusOne(digits: Array[Int]): Array[Int] = {
        
        def inc(d: List[Int]) : List[Int] = d match {
            case Nil => List(1)
            case head :: tail =>
                if (head == 9) 0 :: inc(tail)
                else  (head + 1) :: tail
        }

        val input = digits.reverse.toList
        val output = inc(input).reverse

        output.toArray
    }
}
