package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const StatementRegex = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
const NumbersRegex = `mul\((\d{1,3}),(\d{1,3})\)`

func parseMulStatements(fileName string) []string {
	var statements []string
	re := regexp.MustCompile(StatementRegex)

	inputFile, err := os.Open(fileName)
	if err != nil {
		return statements
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		statements = append(statements, matches...)
	}
	return statements

}

func processMulStatement(statement string) int {
	re := regexp.MustCompile(NumbersRegex)
	match := re.FindStringSubmatch(statement)
	num1, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	return num1 * num2
}

func processStatements(statements []string) int {
	enabled := true
	sum := 0
	for _, statement := range statements {
		if statement == "do()" {
			enabled = true
		} else if statement == "don't()" {
			enabled = false
		} else {
			if enabled {
				sum += processMulStatement(statement)
			}
		}
	}
	return sum
}

func main() {
	statements := parseMulStatements("input.txt")
	sum := processStatements(statements)
	fmt.Println(sum)
}
