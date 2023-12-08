package main

import (
	"bufio"
	"fmt"
	"os"
)

func day8() {
	day8part1()
	day8part2()
}

type Node struct {
	left  string
	right string
}

type Move struct {
	len     int
	current int
	moves   string
}

func (move *Move) nextMove() byte {
	result := move.moves[move.current]
	move.current = (move.current + 1) % move.len
	return result
}

func day8part1() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	moves := rows[0]
	move := Move{
		len:     len(moves),
		current: 0,
		moves:   moves,
	}

	nodes := make(map[string]Node, 0)

	for _, row := range rows[2:] {
		node := row[0:3]
		left := row[7:10]
		right := row[12:15]
		nodes[node] = Node{
			left:  left,
			right: right,
		}
	}

	numMoves := 0
	current := "AAA"
	destination := "ZZZ"

	for current != destination {
		leftOrRight := move.nextMove()
		if leftOrRight == 'L' {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
		numMoves++
	}

	fmt.Println("result:", numMoves)
}

func day8part2() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
	}

	moves := rows[0]

	nodes := make(map[string]Node, 0)
	current := make([]string, 0)

	for _, row := range rows[2:] {
		node := row[0:3]
		left := row[7:10]
		right := row[12:15]
		nodes[node] = Node{
			left:  left,
			right: right,
		}
		if node[2] == 'A' {
			current = append(current, node)
		}
	}

	cycleLengths := make([]int, 0, len(current))

	for i := range current {
		cycleLengths = append(cycleLengths, detectCycleLength(current[i], nodes, moves))
	}

	numMoves := LCM(cycleLengths[0], cycleLengths[1], cycleLengths[2:]...)

	fmt.Println("result:", numMoves)
}

// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func detectCycleLength(start string, nodes map[string]Node, moves string) int {
	slow := start
	fast := start
	slowMover := Move{
		len:     len(moves),
		current: 0,
		moves:   moves,
	}
	fastMover := Move{
		len:     len(moves),
		current: 0,
		moves:   moves,
	}

	if slowMover.nextMove() == 'L' {
		slow = nodes[slow].left
	} else {
		slow = nodes[slow].right
	}
	if fastMover.nextMove() == 'L' {
		fast = nodes[fast].left
	} else {
		fast = nodes[fast].right
	}
	if fastMover.nextMove() == 'L' {
		fast = nodes[fast].left
	} else {
		fast = nodes[fast].right
	}

	for slow != fast {
		if slowMover.nextMove() == 'L' {
			slow = nodes[slow].left
		} else {
			slow = nodes[slow].right
		}
		if fastMover.nextMove() == 'L' {
			fast = nodes[fast].left
		} else {
			fast = nodes[fast].right
		}
		if fastMover.nextMove() == 'L' {
			fast = nodes[fast].left
		} else {
			fast = nodes[fast].right
		}
	}

	cycleMoveStart := slowMover.current
	if slowMover.nextMove() == 'L' {
		slow = nodes[slow].left
	} else {
		slow = nodes[slow].right
	}
	numMoves := 1

	for slow != fast || slowMover.current != cycleMoveStart {
		if slowMover.nextMove() == 'L' {
			slow = nodes[slow].left
		} else {
			slow = nodes[slow].right
		}
		numMoves++
	}

	return numMoves
}
