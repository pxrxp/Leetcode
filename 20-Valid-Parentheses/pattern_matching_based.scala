object Solution {
    def isValid(s: String): Boolean = {

        def pair(c: Char): Char = c match {
            case '(' => ')'
            case '[' => ']'
            case '{' => '}'
        }

        // Push all opening brackets in stack
        // Pop them when their pair closing brackets encountered
        // LIFO order ensures matching brackets
        // If stack empty at the end of processing, is valid.

        @annotation.tailrec
        def validate(chars: List[Char], stack: List[Char] = Nil): Boolean = chars match {
            case Nil => stack.isEmpty

            case char :: rest => char match {
                case '(' | '[' | '{' => validate(rest, char :: stack)
                case ')' | ']' | '}' => stack match {
                    case head :: tail if pair(head) == char => validate(rest, tail)
                    case _ => false // Handles both stack.isEmpty and stack not having pair char
                }
                case _ => false
            }
        }

        validate(s.toList)
    }
}
