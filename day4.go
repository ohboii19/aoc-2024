package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseGridInput(inputFileName string) []string {
	var grid []string
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening file")
		return grid
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getFeasibleDirections(grid []string, word string, pos_y int, pos_x int) [][2]int {
	var directions [][2]int
	wordLen := len(word)
	gridHeight := len(grid)
	gridWidth := len(grid[pos_y])

	// Directions (dx, dy)
	dirOffsets := [][2]int{
		{0, 1}, {0, -1}, {-1, 0}, {1, 0}, // Down, Up, Left, Right
		{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, // Down-right, Up-right, Up-left, Down-left
	}

	for _, dir := range dirOffsets {
		dx, dy := dir[0], dir[1]
		if pos_x+dx*(wordLen-1) >= 0 && pos_x+dx*(wordLen-1) < gridWidth &&
			pos_y+dy*(wordLen-1) >= 0 && pos_y+dy*(wordLen-1) < gridHeight {
			directions = append(directions, dir)
		}
	}

	return directions
}

func getNumWordsFromPosition(grid []string, word string, pos_y int, pos_x int) int {
	count := 0
	wordLen := len(word)

	feasibleDirections := getFeasibleDirections(grid, word, pos_y, pos_x)

	for _, direction := range feasibleDirections {
		dy, dx := direction[1], direction[0]
		cur_y, cur_x := pos_y+dy, pos_x+dx
		for i := 1; i < wordLen; i++ {
			letter := string(grid[cur_y][cur_x])
			if letter != string(word[i]) {
				break
			}
			if i == wordLen-1 {
				count++
			}
			cur_y, cur_x = cur_y+dy, cur_x+dx
		}
	}

	return count
}

func getNumWordOccurences(grid []string, word string) int {
	count := 0
	for row_i := 0; row_i < len(grid); row_i++ {
		for col_j := 0; col_j < len(grid[row_i]); col_j++ {
			if grid[row_i][col_j] == word[0] {
				count += getNumWordsFromPosition(grid, word, row_i, col_j)
			}
		}
	}
	return count
}

func main() {
	grid := parseGridInput("toy_input.txt")
	fmt.Println(getNumWordOccurences(grid, "XMAS"))
	fmt.Println(getNumWordsFromPosition(grid, "MAS", 5, 4))
}

/* Thoughts on part2
Everytime you encounter a MAS, store its centre position in a map, mapping each centre position to a count

when the count is 2, you have 1 cross

we are now only considering diagonal directions
*/
