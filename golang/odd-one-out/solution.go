package solution

// chatgpt said that using xor would be more space efficient
// oddNum := 0
// for _, num := range A {
//      oddNum ^= num  
// } 
// return oddNum
// but this was my solution (which got 100% Task Score, Correctness,
// and Performance):

func Solution(A []int) int {
    isOdd := make(map[int]bool)
    for _, num := range A {
        if isOdd[num] {
            isOdd[num] = false
        } else {
            isOdd[num] = true
        }
    }
    for num, odd := range isOdd {
        if odd {
            return num
        }
    }
    return 0
}

