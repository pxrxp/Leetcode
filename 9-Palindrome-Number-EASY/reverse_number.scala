object Solution {
    def isPalindrome(x: Int): Boolean = {
        def reverse(num: Int, result: Int = 0): Int = {
            if (num == 0) result
            else reverse(num / 10, result * 10 + num % 10)
        }

        x >= 0 && x == reverse(x)
    }
}
