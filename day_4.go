package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day4() {
	day4part1()
	day4part2()
}

func day4part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ":")
		numSets := strings.Split(split[1], "|")
		total += calcCardPoints(strings.TrimSpace(numSets[0]), strings.TrimSpace(numSets[1]))
	}

	fmt.Println("Num:", total)
}

func calcCardPoints(winners string, mine string) int {
	winnersAry := strings.Split(winners, " ")
	winnerMap := make(map[int]struct{})
	for _, winner := range winnersAry {
		winnerInt, _ := strconv.Atoi(winner)
		winnerMap[winnerInt] = struct{}{}
	}

	score := 0
	myAry := strings.Split(mine, " ")
	for _, my := range myAry {
		myInt, er := strconv.Atoi(strings.TrimSpace(my))
		if er != nil {
			continue
		}
		if _, exists := winnerMap[myInt]; exists {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}

func calcWinners(winners string, mine string) int {
	winnersAry := strings.Split(winners, " ")
	winnerMap := make(map[int]struct{})
	for _, winner := range winnersAry {
		winnerInt, _ := strconv.Atoi(winner)
		winnerMap[winnerInt] = struct{}{}
	}

	score := 0
	myAry := strings.Split(mine, " ")
	for _, my := range myAry {
		myInt, er := strconv.Atoi(strings.TrimSpace(my))
		if er != nil {
			continue
		}
		if _, exists := winnerMap[myInt]; exists {
			score++
		}
	}

	return score
}

func day4part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	cardWins := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ":")
		numSets := strings.Split(split[1], "|")
		cardWins = append(cardWins, calcWinners(strings.TrimSpace(numSets[0]), strings.TrimSpace(numSets[1])))
	}

	total := 0
	for i := 0; i < len(cardWins); i++ {
		total += calcTotal(cardWins, i, "")
	}

	fmt.Println("Num:", total)
}

func calcTotal(cardWins []int, cur int, indent string) int {
	wins := cardWins[cur]
	result := 1
	for i := 1; i <= wins; i++ {
		result += calcTotal(cardWins, cur+i, indent+"  ")
	}

	return result
}
