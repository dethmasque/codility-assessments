package solution

func Solution(A []int) int {
    size := len(A)
    seen := make(map[int]bool)
    for _, num := range A {
        if num > 0 {
            seen[num] = true
        }           
    }
    for i := 1; i <= size; i++ {
        if !seen[i] {
            return i
        }
    }
    return size + 1 
}
