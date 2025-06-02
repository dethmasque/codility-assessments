package solution

func Solution(A []int) int {
    unique := 0
    seen := make(map[int]bool)
    for _, num := range A {
        if !seen[num] {
            seen[num] = true
            unique++
        }
    }
    return unique
}

