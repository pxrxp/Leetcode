object Solution {
    def moveZeroes(nums: Array[Int]): Unit = {} {
        nums.indices.foldLeft(0) {
            (nonZeroIndex, i) => nums(i) match {
                case 0 => nonZeroIndex
                case _ =>
                    nums(nonZeroIndex) = nums(i)
                    if (i != nonZeroIndex) nums(i) = 0
                    nonZeroIndex + 1
            }
        }
    }
}
