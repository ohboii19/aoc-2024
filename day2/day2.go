package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseReports(fileName string) ([][]int, error) {
	var reports [][]int

	inputFile, err := os.Open(fileName)
	if err != nil {
		return reports, errors.New("Error opening input file")
	}
	defer inputFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		var report []int
		// Print each line
		line := scanner.Text()
		levelStrings := strings.Fields(line)
		for _, str := range levelStrings {
			if num, err := strconv.Atoi(str); err == nil {
				report = append(report, num)
			} else {
				fmt.Println("Skipping invalid input:", str)
			}
		}
		reports = append(reports, report)

	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return reports, nil
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isReportSafe(report []int) bool {
	isSteady := true
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] {
			isIncreasing = false
		}
		if report[i] <= report[i+1] {
			isDecreasing = false
		}
		if absInt(report[i]-report[i+1]) > 3 {
			isSteady = false
		}
	}

	return (isIncreasing || isDecreasing) && isSteady

}

func isReportDampened(unsafeReport []int) bool {
	for i := 0; i < len(unsafeReport); i++ {
		candidate := append([]int(nil), unsafeReport[:i]...)
		candidate = append(candidate, unsafeReport[i+1:]...)
		if isReportSafe(candidate) {
			return true
		}
	}
	return false
}

func getNumSafeReports(reports [][]int) int {
	numSafeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			numSafeReports++
		} else if isReportDampened(report) {
			numSafeReports++
		}
	}
	return numSafeReports
}

func main() {
	reports, err := parseReports("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Safe reports:", getNumSafeReports(reports))

}
