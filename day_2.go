package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2() {
	day2part1()
	day2part2()
}

func day2part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		colon := strings.Index(text, ":")
		game, _ := strconv.Atoi(text[5:colon])
		if possible(text[colon+1:]) {
			total += game
		}
	}
	fmt.Println("Num:", total)
}

func possible(line string) bool {
	pulls := strings.Split(line, ";")
	for _, pull := range pulls {
		colors := strings.Split(pull, ",")
		for _, color := range colors {
			numColor := strings.Split(strings.TrimSpace(color), " ")
			num, _ := strconv.Atoi(numColor[0])
			color := numColor[1]
			// only 12 red cubes, 13 green cubes, and 14 blue cubes?
			if color == "red" && num > 12 || color == "green" && num > 13 || color == "blue" && num > 14 {
				return false
			}
		}
	}
	return true
}

func power(line string) int {
	pulls := strings.Split(line, ";")
	var mins = map[string]int{
		"red": 0, "green": 0, "blue": 0,
	}
	for _, pull := range pulls {
		colors := strings.Split(pull, ",")
		for _, color := range colors {
			numColor := strings.Split(strings.TrimSpace(color), " ")
			num, _ := strconv.Atoi(numColor[0])
			color := numColor[1]
			if num > mins[color] {
				mins[color] = num
			}
		}
	}
	return mins["red"] * mins["green"] * mins["blue"]
}

func day2part2() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		colon := strings.Index(text, ":")
		total += power(text[colon+1:])

	}
	fmt.Println("Num:", total)
}
