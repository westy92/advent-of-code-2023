package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day5() {
	day5part1()
	day5part2()
}

type MyMap struct {
	source int
	dest   int
	count  int
}

func day5part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	lineNum := 0
	seedsStrs := strings.Split(rows[0][7:], " ")
	seeds := make([]int, 0, len(seedsStrs))
	for _, seed := range seedsStrs {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}

	lineNum += 3
	seedToSoil := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		seedToSoil = append(seedToSoil, parseMyMap(line))
	}

	lineNum += 2
	soil2Fertilizer := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		soil2Fertilizer = append(soil2Fertilizer, parseMyMap(line))
	}

	lineNum += 2
	fertilizer2Water := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		fertilizer2Water = append(fertilizer2Water, parseMyMap(line))
	}

	lineNum += 2
	water2Light := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		water2Light = append(water2Light, parseMyMap(line))
	}

	lineNum += 2
	light2Temp := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		light2Temp = append(light2Temp, parseMyMap(line))
	}

	lineNum += 2
	temp2Humid := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		temp2Humid = append(temp2Humid, parseMyMap(line))
	}

	lineNum += 2
	humid2Loc := make([]*MyMap, 0)
	for lineNum < len(rows) && rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		humid2Loc = append(humid2Loc, parseMyMap(line))
	}

	minLoc := 999999999 // TODO maxInt
	for _, seed := range seeds {
		converted := convertValue(convertValue(convertValue(convertValue(convertValue(convertValue(convertValue(seed, seedToSoil), soil2Fertilizer), fertilizer2Water), water2Light), light2Temp), temp2Humid), humid2Loc)
		if converted < minLoc {
			minLoc = converted
		}
	}

	fmt.Println("min:", minLoc)
}

func convertValue(val int, mapData []*MyMap) int {
	for _, data := range mapData {
		if val >= data.source && val <= data.source+data.count {
			//fmt.Println("convert %v to %v", val, val-data.source+data.dest)
			return val - data.source + data.dest
		}
	}
	//fmt.Println("keep %v", val)
	return val
}

func parseMyMap(line string) *MyMap {
	strs := strings.Split(line, " ")
	destination, _ := strconv.Atoi(strs[0])
	source, _ := strconv.Atoi(strs[1])
	count, _ := strconv.Atoi(strs[2])
	return &MyMap{
		source: source,
		dest:   destination,
		count:  count,
	}
}

func day5part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	lineNum := 0
	seedsStrs := strings.Split(rows[0][7:], " ")
	seeds := make([]int, 0, len(seedsStrs))
	for _, seed := range seedsStrs {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}

	lineNum += 3
	seedToSoil := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		seedToSoil = append(seedToSoil, parseMyMap(line))
	}

	lineNum += 2
	soil2Fertilizer := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		soil2Fertilizer = append(soil2Fertilizer, parseMyMap(line))
	}

	lineNum += 2
	fertilizer2Water := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		fertilizer2Water = append(fertilizer2Water, parseMyMap(line))
	}

	lineNum += 2
	water2Light := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		water2Light = append(water2Light, parseMyMap(line))
	}

	lineNum += 2
	light2Temp := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		light2Temp = append(light2Temp, parseMyMap(line))
	}

	lineNum += 2
	temp2Humid := make([]*MyMap, 0)
	for rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		temp2Humid = append(temp2Humid, parseMyMap(line))
	}

	lineNum += 2
	humid2Loc := make([]*MyMap, 0)
	for lineNum < len(rows) && rows[lineNum] != "" {
		line := rows[lineNum]
		lineNum++
		humid2Loc = append(humid2Loc, parseMyMap(line))
	}

	minLoc := 999999999 // TODO maxInt
	// brute force is GROSS but finds an answer that works in a few minutes...I can do better.
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		len := seeds[i+1]
		end := start + len
		for j := start; j < end; j++ {
			converted := convertValue(convertValue(convertValue(convertValue(convertValue(convertValue(convertValue(j, seedToSoil), soil2Fertilizer), fertilizer2Water), water2Light), light2Temp), temp2Humid), humid2Loc)
			if converted < minLoc {
				fmt.Println("Found smaller. %d", converted)
				minLoc = converted
			}
		}
		fmt.Println("Finished range.")
	}

	fmt.Println("min:", minLoc)
}
