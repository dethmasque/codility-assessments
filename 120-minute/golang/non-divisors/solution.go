package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) []int {
    N := len(A)
    maxVal := 0
    for _, val := range A {
        if val > maxVal {
            // reasonable limit to divisors we compute is the largest int in A
            maxVal = val
        }
    }
    valCount := make([]int, maxVal+1) // much performance
    for _, val := range A {
        valCount[val]++
    }
    divisorCount := make([]int, maxVal+1)
    // assume i is a potential divisor
    for i := 1; i <= maxVal; i++ {
        // calculate all divisors up to the max val
        for j := i; j <= maxVal; j += i {
            // add divisor to all occurences of it
            divisorCount[j] += valCount[i]
        }
    }
    result := make([]int, N)
    for i, val := range A {
        // non-divisors is total numbers in A (N) minus the divisors
        result[i] = N - divisorCount[val]
    }
    return result
}
