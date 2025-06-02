package solution

func Solution(X int, A []int) int {
    size := len(A)
    seen := make(map[int]bool, size)
    goal := X
    uniqueNums := 0 
    for index, num := range A {
        if num >= 1 && num <= X && !seen[num] {
            seen[num] = true
            uniqueNums++
        }
        if uniqueNums == goal {
            return index
        }
    }
    return -1
}

