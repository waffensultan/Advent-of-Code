package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// -- PART 2 STEPS --
	// 1. Count exactly how often each number in the left list appears in the right list
	// --- Using a hashmap (key-value <-- value is for the amount of times it appears)
	// 2. Calculate a Total Similarity Score by:
	// --- Add up each number in the left list AFTER multiplying it by the number of times
	// --- that it appears in the right list

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

	leftMap := make(map[int]int)
	for i := 0; i < len(right); i++ {
		value, error := contains(left, right[i])
		if error == nil {
			leftMap[left[value]] = leftMap[left[value]] + 1
		}
	}

	products := []int{}
	for key, val := range leftMap {
		product := key * val
		products = append(products, product)
	}

	res := 0
	for i := 0; i < len(products); i++ {
		res += products[i]
	}

	fmt.Println(res)
}

/* Returns the index of the duplicate from the passed array. */
func contains(arr []int, val int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i, nil
		}
	}

	return -1, errors.New("value not found")
}
