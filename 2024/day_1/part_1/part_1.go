package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// -- PART 1 STEPS -
	// 1. Save our puzzle input in a file
	// 2. Read the file
	// 3. Save each values in two separate integer arrays (left and right)
	// 4. Sort each list
	// 5. Sum the differences of each index! :)

	data, err := os.ReadFile("../puzzle_input.txt")
	if err != nil {
		panic(err)
	}

	locationCodes := strings.Fields(string(data[:]))

	left := []int{}
	right := []int{}

	for i := 0; i < len(locationCodes); i++ {
		value, err := strconv.Atoi(locationCodes[i])
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	totalDistance := 0.0
	index := 0

	for index < len(right) || index < len(left) {
		leftLocationID := float64(left[index])
		rightLocationID := float64(right[index])

		totalDistance += math.Max(leftLocationID, rightLocationID) - math.Min(leftLocationID, rightLocationID)

		index++
	}

	fmt.Println(int(totalDistance))
}
