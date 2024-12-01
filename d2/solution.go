package d2

import (
	"fmt"
)

/*
TLDR
Find safe reports(gradual dec or inc)
- ie all levels increase OR all levels decrease
- any 2 adjacent levels diff by at least 1,2, or 3
*/
func NumOfSafeReports(reports [][]int) int {
	safeReportsCount := 0
	// iterate reports
	for i := 0; i < len(reports); i++ {
		lc := getLevelChanges(reports[i])
		lcCountPlusEnd := len(lc) + 1

		fmt.Printf("CurrentReport: %d, lc: %s ", reports[i], lc)
		allLevelsSafe := allLevelIncreaseOrDecreasSafely(lc)
		if allLevelsSafe && lcCountPlusEnd == len(reports[i]) {
			fmt.Printf("Safe Report\n\n")
			safeReportsCount++
		} else {
			fmt.Printf("Not a safe report\n\n")
		}

	}

	fmt.Printf("===>Number of Safe reports %d<===\n\n\n", safeReportsCount)
	return safeReportsCount
}

/*
TLDR
Find safe reports(gradual dec or inc)
- ie all levels increase OR all levels decrease
- any 2 adjacent levels diff by at least 1,2, or 3
- for each repor the first instance of a bad level change can be removed

	7 6 4 2 1: Safe without removing any level.
	1 2 7 8 9: Unsafe regardless of which level is removed.
	9 7 6 2 1: Unsafe regardless of which level is removed.
	1 3 2 4 5: Safe by removing the second level, 3.
	8 6 4 4 1: Safe by removing the third level, 4.
	1 3 6 7 9: Safe without removing any level.
*/
func NumOfSafeReportsPart2(reports [][]int) int {
	fmt.Println("=============Part2 WIP==============")
	safeReportsCount := 0
	// iterate reports
	for i := 0; i < len(reports); i++ {
		lc := getLevelChanges(reports[i])
		lcCountPlusEnd := len(lc) + 1

		fmt.Printf("CurrentReport: %d, lc: %s ", reports[i], lc)
		allLevelsSafe := allLevelIncreaseOrDecreasSafely(lc)
		if allLevelsSafe && lcCountPlusEnd == len(reports[i]) {
			fmt.Printf("Safe Report\n\n")
			safeReportsCount++
		} else if hasOnlyOneDeviation(lc) || allLevelsSafe && lcCountPlusEnd == len(reports[i]) {
			fmt.Printf("Safe Report\n\n")
			safeReportsCount++
		} else {
			fmt.Printf("Not a safe report\n\n")
		}

	}

	fmt.Printf("===>Number of Safe reports %d<===\n\n\n", safeReportsCount)
	return safeReportsCount
}

// TODO: yeet and rethink
func hasOnlyOneDeviation(s []string) bool {
	countD := 0
	countI := 0

	for _, v := range s {
		switch v {
		case "d":
			countD++
		case "i":
			countI++
		default:
			return false
		}
	}

	oneDec := countD == 1 && countI == len(s)-1
	oneInc := countI == 1 && countD == len(s)-1

	return oneDec || oneInc
}

// all level change by 1,2, or 3
func allLevelIncreaseOrDecreasSafely(s []string) bool {
	// if empty level change slice then its unsafe, return false
	if len(s) == 0 {
		return false
	}

	// skipOneBadLevel := true

	first := s[0]
	for _, str := range s[1:] {
		if str != first { //&& skipOneBadLevel {
			// skipOneBadLevel = false
			// continue
			// } else {
			return false
		}
	}
	return true
}

func getLevelChanges(r []int) []string {
	levelChanges := []string{}
	for i := 0; i < len(r)-1; i++ {
		c := r[i]
		n := r[i+1]

		// looping reverse so inc/dec is opposite in this case
		change := c - n
		switch change {
		case 1, 2, 3:
			levelChanges = append(levelChanges, "d")
		case -1, -2, -3:
			levelChanges = append(levelChanges, "i")

		default:
			continue
		}
	}

	return levelChanges
}
