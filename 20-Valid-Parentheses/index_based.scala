object Solution {
    def isValid(s: String): Boolean = {

        def pair(c: Char): Char = c match {
            case '(' => ')'
            case '[' => ']'
            case '{' => '}'
        }
        
        @annotation.tailrec
        def validate(s: String, index: Int = 0, stack: List[Char] = Nil): Boolean = {
            if (index >= s.length)
                stack.isEmpty
            else {
                val c: Char = s.charAt(index)
                c match {
                    case '(' | '[' | '{' =>
                        validate(s, index + 1, c :: stack)
                    case ')' | ']' | '}' => {
                        if (stack.nonEmpty && pair(stack.head) == c)
                            validate(s, index + 1, stack.tail)
                        else
                            false
                    }
                }
            }
        }

        validate(s)
    }
}
