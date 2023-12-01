package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func day1() {
	day1part1()
	day1part2()
}

func day1part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()

		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNumber(text), lastNumber(text)))
		total += num
	}
	fmt.Println("Num:", total)
}

func firstNumber(s string) string {
	for _, char := range s {
		if unicode.IsNumber(char) {
			return string(char)
		}
	}
	return ""
}

func lastNumber(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if unicode.IsNumber(c) {
			return string(c)
		}
	}
	return ""
}

var lookup = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func firstNumberWords(s string) string {
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if unicode.IsNumber(c) {
			return string(c)
		}
		if i+2 > len(s) {
			continue
		}
		var word = lookup[s[i:i+2]]
		if word != "" {
			return word
		}
		if i+3 > len(s) {
			continue
		}
		word = lookup[s[i:i+3]]
		if word != "" {
			return word
		}
		if i+4 > len(s) {
			continue
		}
		word = lookup[s[i:i+4]]
		if word != "" {
			return word
		}
		if i+5 > len(s) {
			continue
		}
		word = lookup[s[i:i+5]]
		if word != "" {
			return word
		}
	}
	return ""
}

func lastNumberWords(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if unicode.IsNumber(c) {
			return string(c)
		}
		if i < 2 {
			continue
		}
		var word = lookup[s[i-2:i+1]]
		if word != "" {
			return word
		}
		if i < 3 {
			continue
		}
		word = lookup[s[i-3:i+1]]
		if word != "" {
			return word
		}
		if i < 4 {
			continue
		}
		word = lookup[s[i-4:i+1]]
		if word != "" {
			return word
		}
		if i < 5 {
			continue
		}
		word = lookup[s[i-5:i+1]]
		if word != "" {
			return word
		}
	}
	return ""
}

func day1part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()

		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNumberWords(text), lastNumberWords(text)))
		fmt.Println("", num)
		total += num
	}
	fmt.Println("Num:", total)
}
