object Solution {
    def moveZeroes(nums: Array[Int]): Unit = {
        val nonZeroIndex = moveNonZeroes(nums)
        fillZeroes(nums, nonZeroIndex)
    }

    def moveNonZeroes(nums: Array[Int], nonZeroIndex: Int = 0, index: Int = 0): Int = {
        if (index == nums.length) return nonZeroIndex

        val n = nums(index)

        if (n != 0) {
            nums(nonZeroIndex) = n
            moveNonZeroes(nums, nonZeroIndex + 1, index + 1)
        } else {
            moveNonZeroes(nums, nonZeroIndex, index + 1)
        }
    }

    def fillZeroes(nums: Array[Int], index: Int): Unit = {
        if (index == nums.length) return

        nums(index) = 0
        fillZeroes(nums, index + 1)
    }
}
