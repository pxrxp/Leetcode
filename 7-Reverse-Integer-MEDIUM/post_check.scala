object Solution {
	def reverse(x: Int): Int = {

		@annotation.tailrec
		def recur(num: Int, result: Int = 0): Int = {
			if (num == 0) result
			else {
				val digit = num % 10
				val next = result * 10 + digit
				
        // Check if silently overflowed
				if ((next - digit) / 10 != result) 0
				else recur(num / 10, next)
			}
		}

		recur(x)
	}
}
