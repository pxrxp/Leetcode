object Solution {
    def intersection(nums1: Array[Int], nums2: Array[Int]): Array[Int] = {
        findCommon(nums1.toSet, nums2.toSet).toArray
    }

    def findCommon(set1: Set[Int], set2: Set[Int], result: Set[Int] = Set.empty): Set[Int] = {
        if (set1.isEmpty) result  // Checked every element of set1 if already in set2
        else {
            val head = set1.head
            val tail = set1.tail

            if (set2 contains head) findCommon(tail, set2, result + head)
            else findCommon(tail, set2, result)
        }
    }
}
