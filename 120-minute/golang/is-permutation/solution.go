package solution

func Solution(A []int) int {
    n := len(A)
    seen := make(map[int]bool)
    for i := 0; i < n; i++ {
        num := A[i]
        if seen[num] || num > n {
            return 0
        }
        seen[num] = true
    }
    return 1
}

