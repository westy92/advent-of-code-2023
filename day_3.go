package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"golang.org/x/exp/maps"
)

func day3() {
	day3part1()
	day3part2()
}

type Coordinate struct {
	row int
	col int
}

func day3part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	rows := make([]string, 0)
	startingNumCoords := make(map[Coordinate]struct{}, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			char := rows[i][j]
			if char != '.' && !unicode.IsNumber(rune(char)) {
				// is symbol
				// look at all surrounding areas in below order, if number, add coordinate of first character
				// 123
				// 4 5
				// 678
				// prev row
				if i-1 >= 0 && j-1 >= 0 && unicode.IsNumber(rune(rows[i-1][j-1])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j-1)}] = struct{}{}
				}
				if i-1 >= 0 /*&& j >= 0*/ && unicode.IsNumber(rune(rows[i-1][j])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j)}] = struct{}{}
				}
				if i-1 >= 0 && j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i-1][j+1])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j+1)}] = struct{}{}
				}
				// curr row
				if j-1 >= 0 && unicode.IsNumber(rune(rows[i][j-1])) {
					startingNumCoords[Coordinate{row: i, col: firstNumericDigit(rows[i], j-1)}] = struct{}{}
				}
				if j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i][j+1])) {
					startingNumCoords[Coordinate{row: i, col: firstNumericDigit(rows[i], j+1)}] = struct{}{}
				}
				// next row
				if i+1 < len(rows) && j-1 >= 0 && unicode.IsNumber(rune(rows[i+1][j-1])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j-1)}] = struct{}{}
				}
				if i+1 < len(rows) /*&& j >= 0 */ && unicode.IsNumber(rune(rows[i+1][j])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j)}] = struct{}{}
				}
				if i+1 < len(rows) && j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i+1][j+1])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j+1)}] = struct{}{}
				}
			}
		}
	}

	for k, _ := range startingNumCoords {
		start := k.col
		for start+1 < len(rows[k.row]) && unicode.IsNumber(rune(rows[k.row][start+1])) {
			start++
		}
		num, _ := strconv.Atoi(rows[k.row][k.col : start+1])
		total += num
	}

	fmt.Println("Num:", total)
}

func firstNumericDigit(row string, start int) int {
	result := start
	for result-1 >= 0 && unicode.IsNumber(rune(row[result-1])) {
		result--
	}
	return result
}
func day3part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	rows := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			char := rows[i][j]
			if char == '*' {
				startingNumCoords := make(map[Coordinate]struct{}, 0)
				// is gear
				// look at all surrounding areas in below order, if number, add coordinate of first character
				// 123
				// 4 5
				// 678
				// prev row
				if i-1 >= 0 && j-1 >= 0 && unicode.IsNumber(rune(rows[i-1][j-1])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j-1)}] = struct{}{}
				}
				if i-1 >= 0 /*&& j >= 0*/ && unicode.IsNumber(rune(rows[i-1][j])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j)}] = struct{}{}
				}
				if i-1 >= 0 && j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i-1][j+1])) {
					startingNumCoords[Coordinate{row: i - 1, col: firstNumericDigit(rows[i-1], j+1)}] = struct{}{}
				}
				// curr row
				if j-1 >= 0 && unicode.IsNumber(rune(rows[i][j-1])) {
					startingNumCoords[Coordinate{row: i, col: firstNumericDigit(rows[i], j-1)}] = struct{}{}
				}
				if j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i][j+1])) {
					startingNumCoords[Coordinate{row: i, col: firstNumericDigit(rows[i], j+1)}] = struct{}{}
				}
				// next row
				if i+1 < len(rows) && j-1 >= 0 && unicode.IsNumber(rune(rows[i+1][j-1])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j-1)}] = struct{}{}
				}
				if i+1 < len(rows) /*&& j >= 0 */ && unicode.IsNumber(rune(rows[i+1][j])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j)}] = struct{}{}
				}
				if i+1 < len(rows) && j+1 < len(rows[i-1]) && unicode.IsNumber(rune(rows[i+1][j+1])) {
					startingNumCoords[Coordinate{row: i + 1, col: firstNumericDigit(rows[i+1], j+1)}] = struct{}{}
				}

				if len(startingNumCoords) == 2 {
					coords := maps.Keys(startingNumCoords)
					total += getNum(rows[coords[0].row], coords[0].col) * getNum(rows[coords[1].row], coords[1].col)
				}
			}
		}
	}

	fmt.Println("Num:", total)
}

func getNum(row string, start int) int {
	cur := start
	for cur+1 < len(row) && unicode.IsNumber(rune(row[cur+1])) {
		cur++
	}
	num, _ := strconv.Atoi(row[start : cur+1])
	return num
}
