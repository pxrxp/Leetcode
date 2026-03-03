object Solution {
    def twoSum(nums: Array[Int], target: Int): Array[Int] = {

        @annotation.tailrec
        def helper(index: Int = 0, seen: Map[Int, Int] = Map.empty): Array[Int] = {
            if (index >= nums.length) Array()
            else {
                val complement = target - nums(index)
                seen.get(complement) match {
                    case Some(complementIndex) => Array(index, complementIndex)
                    case None => helper(index+1, seen+(nums(index)->index))
                }
            }
        }

        helper()
    }
}
