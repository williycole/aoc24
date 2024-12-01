package d1

import (
	"fmt"
	"sort"
)

/*
TLDR

 1. Parse and sort two lists of numbers.
 2. Calculate the absolute difference between corresponding numbers in the sorted lists.
 3. Sum up all the differences to get the total distance.
*/
func TotalDistance(l1, l2 []int) int {
	// Sort Lists
	sort.Ints(l1)
	sort.Ints(l2)

	// we want the abosolute difference so start with larger number
	distances := []int{}
	for i := 0; i < len(l1); i++ {
		if l1[i] >= l2[i] {
			d := l1[i] - l2[i]
			distances = append(distances, d)
		} else {
			d := l2[i] - l1[i]
			distances = append(distances, d)
		}
	}

	// sum, print, return
	td := 0
	for _, d := range distances {
		td += d
	}

	fmt.Printf("Total Distances %d\n", td)
	return td
}

/*
TLDR for Part Two:

 1. Count occurrences of each number from the left list in the right list.
 2. Multiply each number in the left list by its count in the right list.
 3. Sum up all these products to get the total similarity score.
*/
func TotalDistancePart2(l1, l2 []int) int {
	// Sort Lists
	sort.Ints(l1)
	sort.Ints(l2)

	// get left lists item count from right list
	// multiple current item by its count and store for later
	lcounts := []int{}
	for i := 0; i < len(l1); i++ {
		icnt := 0

		for j := 0; j < len(l2); j++ {
			if l1[i] == l2[j] {
				icnt++
			}
		}

		scount := l1[i] * icnt
		lcounts = append(lcounts, scount)

		icnt = 0
	}

	// sum, print, return
	td := 0
	for _, d := range lcounts {
		td += d
	}

	fmt.Printf("Similarity Score %d\n", td)
	return td
}
