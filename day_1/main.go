package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// this is mostly to avoid a return that relies on position
type ListData struct {
	firstList  []int
	secondList []int
}

func readList() ListData {
	file, err := os.Open("./values.txt")

	// We don't want to try and use file if there is an error
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var firstList, secondList []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			firstList = append(firstList, num1)
			secondList = append(secondList, num2)
		}
	}

	lists := ListData{}

	lists.firstList = firstList
	lists.secondList = secondList

	return lists
}

func main() {

	lists := readList()

	// The make function allocates a zeroed array and returns a slice that refers to that array
	emptyList := make([]int, len(lists.firstList))

	sort.Ints(lists.firstList[:])
	sort.Ints(lists.secondList[:])

	for i := 0; i < len(emptyList); i++ {
		var first = lists.firstList[i]
		var second = lists.secondList[i]
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

	println(sum)
}
