object Solution {
    def maxSubArray(nums: Array[Int]): Int = {
        nums.foldLeft((0,Int.MinValue)) {
            case ((currSum, maxSum), elem) =>
                val newCurrSum = Math.max(elem, currSum + elem)
                val newMaxSum = Math.max(maxSum, newCurrSum)
                (newCurrSum, newMaxSum)
        }._2
    }
}
