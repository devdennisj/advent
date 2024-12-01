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
	// ScanLines is a split function for a Scanner that returns each line of text, stripped of any trailing end-of-line marker
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) == 2 {
			// Ataoi is string to int
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

func getDistance(first, second int) int {
	if first == second {
		return 0
	}

	if first > second {
		return first - second
	}

	return second - first

}

func main() {

	/*
		Short Variable Declaration Operator(:=) in Golang is used to create the variables having a proper name and initial value. The main purpose of using this operator to declare and initialize the local variables inside the functions and to narrowing the scope of the variables.
	*/
	lists := readList()

	// The make function allocates a zeroed array and returns a slice that refers to that array
	emptyList := make([]int, len(lists.firstList))

	sort.Ints(lists.firstList)
	sort.Ints(lists.secondList)

	for i := 0; i < len(emptyList); i++ {
		var first = lists.firstList[i]
		var second = lists.secondList[i]

		emptyList[i] = getDistance(first, second)
	}

	var sum = 0

	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	for _, v := range emptyList {
		sum += v
	}

	println(sum)
}
