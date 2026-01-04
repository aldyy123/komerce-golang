package testcase

import (
	"fmt"
	"sort"
)

func PSBB(n int, arr []int) int {
	buses := 0

	if n != len(arr) {
		fmt.Println("Input must be equal with count of family")
		return buses
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	used := make([]bool, n)

	for i := 0; i < len(arr); i++ {
		buses++
		used[i] = true
		remaining := 4 - arr[i]
		if remaining > 0 {
			for j := i + 1; j < n; j++ {
				if !used[j] && arr[j] <= remaining {
					used[j] = true
					break
				}
			}
		}
	}

	return buses
}
