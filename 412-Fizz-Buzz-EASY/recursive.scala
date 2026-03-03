object Solution {
    def fizzBuzz(n: Int): List[String] = {
        def fB(i: Int): String = i match {
            case i if (i % 15 == 0) => "FizzBuzz"
            case i if (i % 3 == 0)  => "Fizz"
            case i if (i % 5 == 0)  => "Buzz"
            case i                  => i.toString
        }

        @annotation.tailrec
        def iter(i: Int, acc: List[String] = Nil): List[String] = {
            if (i >= 1) iter(i-1, fB(i) :: acc)
            else acc
        }

        iter(n, Nil)
    }
}
