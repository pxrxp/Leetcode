object Solution {
    def mySqrt(x: Int): Int = {

        @annotation.tailrec
        def binarySearch(low: Int, high: Int): Int = {
            if (low > high) high
            else {
                val mid = low + (high - low)/2
                val square = mid.toLong * mid.toLong

                if (square == x) mid
                else if (square < x) binarySearch(mid+1, high)
                else binarySearch(low, mid-1)
            }
        }

        binarySearch(0, x)
    }
}
