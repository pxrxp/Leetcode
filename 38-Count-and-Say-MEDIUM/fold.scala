object Solution {

	def countAndSay(n: Int): String = {
		if (n == 1) "1"
		else {
			val prev = countAndSay(n - 1)
			describe(prev)
		}
	}

	// Run length encoding
	def describe(s: String): String = {
		s.foldLeft(("", 0, s.head)) {
			case ((result, count, prevChar), currChar) =>
				if (currChar == prevChar)
					(result, count + 1, prevChar)
				else
					(result + count.toString + prevChar, 1, currChar)
		} match {
		  // Handle last group converting tuple accumulator to string
			case (result, count, prevChar) =>
				result + count.toString + prevChar
		}
	}
}
