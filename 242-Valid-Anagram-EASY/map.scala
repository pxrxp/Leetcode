object Solution {
    def isAnagram(s: String, t: String): Boolean = {
        countChars(s.toList, Map.empty) == countChars(t.toList, Map.empty)
    }

    def countChars(chars: List[Char], counts: Map[Char, Int]): Map[Char, Int] = chars match {
        case Nil => counts
        case head :: tail =>
            val newCount = counts.getOrElse(head, 0)+1
            countChars(tail, counts.updated(head, newCount))
    }
}
