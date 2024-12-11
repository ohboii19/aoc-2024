package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
	"sort"
)

func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func parseLists(fileName string) ([]int, []int, error) {
	var nums1 []int
	var nums2 []int

	inputFile, err := os.Open("input.txt")
	if err != nil {
		return nums1, nums2, errors.New("Error opening input file")
	}
	defer inputFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// Print each line
		line := scanner.Text()
		nums :=  strings.Fields(line)
		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])

		if err1 != nil || err2 != nil  {
			return nums1, nums2, errors.New("Error converting nums")
		}
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}

	
	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	return nums1, nums2, nil
}

func calcTotalDistance(nums1 []int, nums2 []int) (int) {
	totalDistance := 0

	listsLength := len(nums1)
	if listsLength != len(nums2) {
		fmt.Println("arrays are not the same length")
		return totalDistance
	}

	for i:=0; i < listsLength; i++ {
		totalDistance += absInt(nums1[i] - nums2[i])
	}

	return totalDistance
}

func calcSimilarityScore(nums1 []int, nums2 []int) (int) {
	similarityScore := 0
	countMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		countMap[nums2[i]]++
	}
	for i := 0; i < len(nums1); i++ {
		similarityScore += nums1[i] * countMap[nums1[i]]
	}

	return similarityScore
}

func main() {
	nums1, nums2, err := parseLists("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Total Distance:", calcTotalDistance(nums1, nums2))
	fmt.Println("Total Distance:", calcSimilarityScore(nums1, nums2))
}