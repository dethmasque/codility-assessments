package solution

import "strconv"

func Solution(N int) int {
    binary := strconv.FormatInt(int64(N), 2)
    maxGap := 0
    gap := -1

    for _, bit := range binary {
        if bit == '1' {
            if gap > maxGap {
                maxGap = gap
            }
            gap = 0
        } else if gap != -1 {
            gap++
        }
    }
    return maxGap
}

