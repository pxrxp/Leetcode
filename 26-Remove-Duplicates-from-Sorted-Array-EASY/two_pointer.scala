object Solution {
    def removeDuplicates(nums: Array[Int]): Int = {
        
        @annotation.tailrec
        def moveUniques(nums: Array[Int], uniqueIndex: Int = 1, index: Int = 0): Int = {
            if (index == nums.length || uniqueIndex == nums.length)
                uniqueIndex

            // If unique element encountered (curr != prev)
            
            if (index > 0 && nums(index) != nums(index-1)) {
                // Put it in the desired index in front (uniqueIndex)
                // Increase the uniqueIndex for next unique number

                nums(uniqueIndex) = nums(index)
                moveUniques(nums, uniqueIndex + 1, index + 1)
            } else {
                // Only move the index pointer

                moveUniques(nums, uniqueIndex, index + 1)
            }
        }

        moveUniques(nums)

    }
}
