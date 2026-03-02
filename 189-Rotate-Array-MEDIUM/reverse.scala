object Solution {
    def rotate(nums: Array[Int], _k: Int): Unit = {
        val n: Int = nums.length
        val k: Int = _k % n

        // Example
        // [1,2,3,4,5,6,7], k = 3
        // ==> [5,6,7,1,2,3,4]

        // [5,6,7] -> [7,6,5], [1,2,3,4] -> [4,3,2,1]
        // [7,6,5,4,3,2,1]


        // PROOF

        // Rotate array by k steps
        // i -> (i+k)%n
        // If i+k>n,  i -> i+k-n (will be on left of k)
        // If i+k<=n, i -> i+k (will be on right of k)

        // is equivalent to

        // Reverse whole array
        // 0 -> n-1, n-1 -> 0
        // i -> (n-1)-i

        // Reverse left of array
        // 0 -> k-1, k-1 -> 0
        // i -> (k-1)-i
        //    = (k-1)-[ (n-1)-i ]
        //    = i+k-n

        // Reverse right of array
        // k -> n-1, n-1 -> k
        // i -> k+(n-1)-i
        //    = k+(n-1)-[ (n-1)-i ]
        //    = i+k

        reverse(nums, 0, n-1)
        reverse(nums, 0, k-1)
        reverse(nums, k, n-1)
    }

    def reverse(nums: Array[Int], start: Int, end: Int): Unit = {
        if (start >= end) return
        swap(nums, start, end)
        reverse(nums, start+1, end-1)
    }

    // Swap using XOR
    def swap(arr: Array[Int], i: Int, j: Int): Unit = {
        arr(i) = arr(i) ^ arr(j) // x^y
        arr(j) = arr(i) ^ arr(j) // (x^y)^y = x^0 = x
        arr(i) = arr(i) ^ arr(j) // (x^y)^x = y^0 = y
    }
}
