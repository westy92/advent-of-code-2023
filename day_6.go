package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day6() {
	day6part1()
	day6part2()
}

type Race struct {
	time     int
	distance int
}

func day6part1() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	timeStrs := strings.Fields(rows[0][5:])
	distanceStrs := strings.Fields(rows[1][9:])

	result := 1
	for i := range timeStrs {
		time, _ := strconv.Atoi(timeStrs[i])
		distance, _ := strconv.Atoi(distanceStrs[i])
		waysToWin := calculateWaysToWin(Race{time: time, distance: distance})
		result *= waysToWin
	}

	fmt.Println("result:", result)
}

func calculateWaysToWin(race Race) int {
	count := 0
	for wait := 0; wait <= race.time; wait++ {
		distance := (race.time - wait) * wait
		if distance > race.distance {
			count++
		}
	}
	return count
}

func day6part2() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	time, _ := strconv.Atoi(strings.ReplaceAll(rows[0][5:], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(rows[1][9:], " ", ""))
	waysToWin := calculateWaysToWin(Race{time: time, distance: distance})

	fmt.Println("result:", waysToWin)
}
