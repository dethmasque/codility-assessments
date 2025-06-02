package solution

import "math"

func Solution(A []int) int {
    n := len(A)
    // 0 < P < n
    // (A[0] + ... + A[P-1]) - (A[P] + A[P + 1] ... + A[n -1])
    // find the minimal difference
    // 1 < n <= 100,000
    // -1000 < A[i] < 1000
    total := 0
    for _, num := range A {
        total += num
    }

    bestDiff := 100000 // reasonable expected max?
    left := 0

    for p := 0; p < n - 1; p++ {
        left += A[p]
        right := total - left
        diff := int(math.Abs(float64(right - left))) // avoid negative values, etc
        if diff < bestDiff {
            bestDiff = diff
        }
    }

    return bestDiff 
}
