package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day9() {
	day9part1()
	day9part2()
}

func day9part1() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numberStrings := strings.Split(line, " ")
		numbers := make([]int, 0, len(numberStrings))
		for _, numStr := range numberStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		next := nextNum(numbers) + numbers[len(numbers)-1]
		fmt.Println(next)
		result += next
	}

	fmt.Println("result:", result)
}

func nextNum(nums []int) int {
	if len(nums) == 1 {
		return 0 // TODO?
	}

	diffs := make([]int, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	firstDiff := diffs[0]
	for _, diff := range diffs[1:] {
		if diff != firstDiff {
			return diffs[len(diffs)-1] + nextNum(diffs)
		}
	}
	return firstDiff
}

func day9part2() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numberStrings := strings.Split(line, " ")
		numbers := make([]int, 0, len(numberStrings))
		for _, numStr := range numberStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		slices.Reverse(numbers) // the only difference haha
		next := nextNum(numbers) + numbers[len(numbers)-1]
		fmt.Println(next)
		result += next
	}

	fmt.Println("result:", result)
}
