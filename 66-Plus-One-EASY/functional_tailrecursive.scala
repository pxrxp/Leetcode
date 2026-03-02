object Solution {
    def plusOne(digits: Array[Int]): Array[Int] = {
        
        @annotation.tailrec
        def inc(d: List[Int], carry: Int, acc: List[Int]) : List[Int] = d match {
            case Nil => 
                if (carry == 0) acc
                else carry :: acc
            case head :: tail =>
                val sum = head + carry
                inc(tail, sum / 10, (sum % 10) :: acc)
        }

        val input = digits.reverse.toList
        val output = inc(input, 1, Nil)

        output.toArray
    }
}
