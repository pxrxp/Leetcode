object Solution {
    def singleNumber(nums: Array[Int]): Int = {
        nums.reduce(_^_)
        // x^x^y^z^y = 0^0^z = z
    }
}
