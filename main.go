package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"main.go/d1"
	d1puzzle "main.go/d1/puzzle-input"
	"main.go/d2"
	d2puzzle "main.go/d2/puzzle-input"
)

func main() {
	flag.Parse()
	// prompt for which puzzle to use and handle shoice
	fmt.Println("Enter your choice (1-25):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 25.")
		return
	}

	switch choice {
	case 1:
		fmt.Println("Test:")
		d1.TotalDistance(d1puzzle.ListThree, d1puzzle.ListFour)
		d1.TotalDistancePart2(d1puzzle.ListThree, d1puzzle.ListFour)
		fmt.Println("\n----------------------------")
		fmt.Println("Answer:")
		d1.TotalDistance(d1puzzle.ListOne, d1puzzle.ListTwo)
		d1.TotalDistancePart2(d1puzzle.ListOne, d1puzzle.ListTwo)
		fmt.Println("\n============================")
	case 2:
		// fmt.Println("Test:")
		fmt.Println(d2puzzle.TestReports)
		// d2.NumOfSafeReports(d2puzzle.TestReports)
		d2.NumOfSafeReportsPart2(d2puzzle.TestReports)
		fmt.Println("\n----------------------------")
		fmt.Println("Answer:")
		// d2.NumOfSafeReports(d2puzzle.Reports)
		// d2.NumOfSafeReportsPart2(d2puzzle.Reports)
		fmt.Println("\n============================")
	case 3:
		fmt.Println("Day 3: TODO")
	case 4:
		fmt.Println("Day 4: TODO")
	case 5:
		fmt.Println("Day 5: TODO")
	case 6:
		fmt.Println("Day 6: TODO")
	case 7:
		fmt.Println("Day 7: TODO")
	case 8:
		fmt.Println("Day 8: TODO")
	case 9:
		fmt.Println("Day 9: TODO")
	case 10:
		fmt.Println("Day 10: TODO")
	case 11:
		fmt.Println("Day 11: TODO")
	case 12:
		fmt.Println("Day 12: TODO")
	case 13:
		fmt.Println("Day 13: TODO")
	case 14:
		fmt.Println("Day 14: TODO")
	case 15:
		fmt.Println("Day 15: TODO")
	case 16:
		fmt.Println("Day 16: TODO")
	case 17:
		fmt.Println("Day 17: TODO")
	case 18:
		fmt.Println("Day 18: TODO")
	case 19:
		fmt.Println("Day 19: TODO")
	case 20:
		fmt.Println("Day 20: TODO")
	case 21:
		fmt.Println("Day 21: TODO")
	case 22:
		fmt.Println("Day 22: TODO")
	case 23:
		fmt.Println("Day 23: TODO")
	case 24:
		fmt.Println("Day 24: TODO")
	case 25:
		fmt.Println("Day 25: TODO")
	default:
		fmt.Println("Please provide a correct puzzle day to run (1-25)")
	}
}
