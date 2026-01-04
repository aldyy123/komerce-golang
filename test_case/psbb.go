package testcase

import (
	"fmt"
	"sort"
)

func PSBB(n int, arr []int) int {

	if n != len(arr) {
		fmt.Print("Input must be equal with count of family")
		return -1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	buses := 0
	used := make([]bool, n)

	for i := 0; i < len(arr); i++ {
		if used[i] {
			continue
		}

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
