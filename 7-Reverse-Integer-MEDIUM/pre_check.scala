object Solution {
	def reverse(x: Int): Int = {
		var num = x
		var result = 0
		while (num != 0) {
			val digit = num % 10

      // Int32: -2,147,483,648 to 2,147,483,647
			
			if (result > Int.MaxValue / 10 || (result == Int.MaxValue / 10 && digit > 7)) return 0
			if (result < Int.MinValue / 10 || (result == Int.MinValue / 10 && digit < -8)) return 0
			result = result * 10 + digit
			num /= 10
		}
		result
	}
}
