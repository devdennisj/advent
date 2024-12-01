package main

import (
	"sort"
)

func main() {
	var emptyList = [6]int{}

	var firstList = [6]int{3, 4, 2, 1, 3, 3}
	sort.Ints(firstList[:])

	var secondList = [6]int{4, 3, 5, 3, 9, 3}
	sort.Ints(secondList[:])

	for i := 0; i < len(emptyList); i++ {
		var first = firstList[i]
		var second = secondList[i]
		if first == second {
			emptyList[i] = 0
		}

		if first > second {
			emptyList[i] = first - second
		} else {
			emptyList[i] = second - first
		}
	}

	var sum = 0

	for _, v := range emptyList {
		sum += v
	}
}
