package solution

func Solution(S string) int {
    size := len(S)
    openParens := make([]rune, 0, size)
    pairs := map[rune]rune{ // map closed to open for match comparison
        ')': '(',
        '}': '{',
        ']': '[',
    }
    for _, paren := range S {
        switch paren {
        case '(', '{', '[':
            openParens = append(openParens, paren)
        case ')', '}', ']':
            if len(openParens) == 0 || openParens[len(openParens)-1] != pairs[paren] {
                return 0
            }
            openParens = openParens[:len(openParens)-1]
        default:
            return 0 // invalid character
        }
    }
    if len(openParens) == 0 {
        return 1
    }
    return 0
}

