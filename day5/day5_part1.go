package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ruleExp = `^\d+\|\d+`
const updateExp = `^\d+(,\d+)*`

func printTable(rulesTable map[string][]string) {
	for key, value := range rulesTable {
		fmt.Println(key, ":", value)
	}
}

func printUpdates(updates [][]string) {
	for _, update := range updates {
		fmt.Println(update)
	}
}

func parseRule(rule string, rulesTable map[string][]string) {
	pageNums := strings.Split(rule, "|")
	rulesTable[pageNums[0]] = append(rulesTable[pageNums[0]], pageNums[1])
}

func parseUpdate(update string, updates *[][]string) {
	pages := strings.Split(update, ",")
	*updates = append(*updates, pages)
}

func parseInputFile(fileName string) (map[string][]string, [][]string) {
	rulesTable := make(map[string][]string)
	var updates [][]string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return rulesTable, updates
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rulesRe := regexp.MustCompile(ruleExp)
		updateRe := regexp.MustCompile(updateExp)
		inputStr := scanner.Text()

		if rulesRe.MatchString(inputStr) {
			parseRule(inputStr, rulesTable)
		} else if updateRe.MatchString(inputStr) {
			parseUpdate(inputStr, &updates)
		}
	}

	return rulesTable, updates
}

func createPageNumToIndexInUpdate(update []string) map[string]int {
	pageNumToIndexInUpdate := make(map[string]int)

	for i, pageNum := range update {
		pageNumToIndexInUpdate[pageNum] = i
	}
	return pageNumToIndexInUpdate
}

func isUpdateValid(rulesTable map[string][]string, update []string) bool {
	pageNumToIndexInUpdate := createPageNumToIndexInUpdate(update)

	for currentIndex, pageNum := range update {
		applicableRules := rulesTable[pageNum]

		for _, rulePageNum := range applicableRules {
			// Check if rulePageNum is in current update
			ruleIndex, ok := pageNumToIndexInUpdate[rulePageNum]
			if ok {
				// If present, check if the index of the rulePageNum is greater than currentIndex
				if currentIndex > ruleIndex {
					return false
				}
			}
		}
	}
	return true
}

func getValidUpdates(rulesTable map[string][]string, updates [][]string) [][]string {
	var validUpdates [][]string

	for _, update := range updates {
		if isUpdateValid(rulesTable, update) {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func sumMiddlePages(validUpdates [][]string) int {
	sum := 0
	for _, update := range validUpdates {
		middlePageNum := update[len(update)/2]
		middlePageNumInt, err := strconv.Atoi(middlePageNum)
		if err != nil {
			fmt.Println("Error coverting to int")
			return -1
		}
		sum += middlePageNumInt
	}
	return sum
}

func main() {
	rulesTable, updates := parseInputFile("toy_input.txt")
	printTable(rulesTable)
	validUpdates := getValidUpdates(rulesTable, updates)
	printUpdates(validUpdates)
	fmt.Println("Sum of the middle page numbers of correct updates:", sumMiddlePages(validUpdates))

}
