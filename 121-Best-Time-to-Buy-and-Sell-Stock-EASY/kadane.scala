object Solution {
	def maxProfit(prices: Array[Int]): Int = {
		prices.foldLeft((Int.MaxValue, 0)) { case ((minPrice, maxProfit), price) =>
			val newMin = Math.min(minPrice, price)
			val newProfit = Math.max(maxProfit, price - minPrice)
			(newMin, newProfit)
		}._2
	}
}
