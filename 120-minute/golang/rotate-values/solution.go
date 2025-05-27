package solution

func Solution(A []int, K int) []int {
    size := len(A)
    if size == 0 {
        return A
    }
    bump := K % size // account for K > size (wrap around)
    if bump != 0 {
        A = append(A[size-bump:], A[:size-bump]...)
    }
    return A
}

