package main

import (
	"bufio"
	"fmt"
	"os"
)

func day10() {
	p1 := day10part1()
	day10part2(p1)
}

type Pipe struct {
	up       bool
	down     bool
	left     bool
	right    bool
	distance int
}

func day10part1() [][]*Pipe {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rowStrings := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rowStrings = append(rowStrings, text)
	}

	height := len(rowStrings)
	width := len(rowStrings[0])

	matrix := make([][]*Pipe, height)
	rows := make([]*Pipe, width*height)
	for i := 0; i < height; i++ {
		matrix[i] = rows[i*width : (i+1)*width]
	}

	/*
			| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
		. is ground; there is no pipe in this tile.
		S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.*/

	start := Coordinate{}
	for row, rowString := range rowStrings {
		for col, colChar := range rowString {
			switch colChar {
			case '|':
				matrix[row][col] = &Pipe{
					up:   true,
					down: true,
				}
			case '-':
				matrix[row][col] = &Pipe{
					left:  true,
					right: true,
				}
			case 'L':
				matrix[row][col] = &Pipe{
					up:    true,
					right: true,
				}
			case 'J':
				matrix[row][col] = &Pipe{
					up:   true,
					left: true,
				}
			case '7':
				matrix[row][col] = &Pipe{
					down: true,
					left: true,
				}
			case 'F':
				matrix[row][col] = &Pipe{
					down:  true,
					right: true,
				}
			case '.':
				matrix[row][col] = nil
			case 'S':
				start = Coordinate{row, col}
				matrix[row][col] = &Pipe{
					down:  true,
					right: true,
					up:    true,
					left:  true,
				}
			}
		}
	}

	fmt.Println(start)

	result := pathLength(matrix, start.row, start.col, 0)

	fmt.Println("result:", result/2)

	return matrix
}

func pathLength(matrix [][]*Pipe, row int, col int, dist int) int {
	// if not a pipe or visited
	cur := matrix[row][col]
	if cur == nil {
		return 0
	} else if cur.distance != 0 {
		return cur.distance
	}

	cur.distance = dist

	connected := make([]Coordinate, 0, 2)
	if cur.up && row-1 >= 0 && matrix[row-1][col] != nil && matrix[row-1][col].down {
		connected = append(connected, Coordinate{row: row - 1, col: col})
	}
	if cur.down && row+1 < len(matrix) && matrix[row+1][col] != nil && matrix[row+1][col].up {
		connected = append(connected, Coordinate{row: row + 1, col: col})
	}
	if cur.left && col-1 >= 0 && matrix[row][col-1] != nil && matrix[row][col-1].right {
		connected = append(connected, Coordinate{row: row, col: col - 1})
	}
	if cur.right && col+1 < len(matrix[0]) && matrix[row][col+1] != nil && matrix[row][col+1].left {
		connected = append(connected, Coordinate{row: row, col: col + 1})
	}

	return max(pathLength(matrix, connected[0].row, connected[0].col, dist+1), pathLength(matrix, connected[1].row, connected[1].col, dist+1))
}

func day10part2(matrix [][]*Pipe) {
	result := 0
	//inside := false
	for row := 0; row < len(matrix); row++ {
		//inside = false
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] != nil {
				fmt.Printf("%3d ", matrix[row][col].distance)
			} else {
				fmt.Print("    ")
			}

			/*if matrix[row][col] != nil && matrix[row][col].distance != 0 {
				fmt.Print("x")
				inside = !inside
			} else {
				fmt.Print(".")
				if inside {
					result++
				}
			}*/

		}
		println("")
	}

	fmt.Println("result:", result)
}
