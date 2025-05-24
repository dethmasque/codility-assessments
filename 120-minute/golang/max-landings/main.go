package main

import (
	"fmt"
	"sort"
)

func MaxLandings(min, max []int) int {
	if len(min) != len(max) {
		return 0
	}

	type landingWindow struct {
		start    int
		end      int
		interval int
	}

	windows := make([]landingWindow, len(min))
	for i := range min {
		windows[i] = landingWindow{min[i], max[i], max[i] - min[i]}
	}

	// give planes who can only land within a short window priority
	sort.Slice(windows, func(i, j int) bool {
		return windows[i].interval < windows[j].interval
	})

	landings := 0
	occupied := make(map[int]bool)
	for index, window := range windows {
		fmt.Printf("Checking plane %d with start %d and end %d \n", index, window.start, window.end)
		// go through the start to the end looking for a free spot
		for i := window.start; i <= window.end; i++ {
			if !occupied[i] {
				fmt.Printf("Found free landing time at %d \n", i)
				occupied[i] = true
				landings++
				break
			}
		}
	}
	fmt.Printf("Number of planes that can land safely: %d \n", landings)

	return landings
}

func main() {
	// test data
	A := []int{1, 2, 4, 8, 0, 10, 0, 2, 4}
	B := []int{3, 7, 5, 15, 15, 12, 1, 5, 4}
	MaxLandings(A, B)
}
