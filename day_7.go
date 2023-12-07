package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func day7() {
	day7part1()
	day7part2()
}

// larger number == stronger
var CamelCardStrength = map[rune]rune{
	'2': 'A',
	'3': 'B',
	'4': 'C',
	'5': 'D',
	'6': 'E',
	'7': 'F',
	'8': 'G',
	'9': 'H',
	'T': 'I',
	'J': 'J',
	'Q': 'K',
	'K': 'L',
	'A': 'M',
}

var CamelCardStrength2 = map[rune]rune{
	'J': 'J', // 74

	'2': 'a', // 97
	'3': 'b',
	'4': 'c',
	'5': 'd',
	'6': 'e',
	'7': 'f',
	'8': 'g',
	'9': 'h',
	'T': 'i',
	'Q': 'k',
	'K': 'l',
	'A': 'm', // 109
}

/* strength order
- 5 of kind (G)
- 4 of kind (F)
- Full House(E)
- 3 of kind (D)
- 2 pair    (C)
- 1 pair    (B)
- high card (A)
*/

type Hand struct {
	Cards []rune
	Bid   int
}

func (hand Hand) priority() string {
	stackedCards := make(map[rune]int, 0)
	cardRanks := make([]rune, 0, len(hand.Cards))

	for _, card := range hand.Cards {
		if _, exists := stackedCards[card]; !exists {
			stackedCards[card] = 1
		} else {
			stackedCards[card]++
		}
		cardRanks = append(cardRanks, CamelCardStrength[card])
	}
	highestCount := 1
	for _, count := range stackedCards {
		if count > highestCount {
			highestCount = count
		}
	}
	multiplier := 'A'

	if highestCount == 5 {
		multiplier = 'G' // 5 of a kind
	} else if highestCount == 4 {
		multiplier = 'F' // 4 of a kind
	} else if highestCount == 3 {
		if len(stackedCards) == 2 {
			multiplier = 'E' // full house
		} else {
			multiplier = 'D' // 3 of a kind
		}
	} else if highestCount == 2 {
		if len(stackedCards) == 3 {
			multiplier = 'C' // 2 pair
		} else {
			multiplier = 'B' // 1 pair
		}
	} else {
		multiplier = 'A' // high card
	}

	return string(multiplier) + string(cardRanks)
}

func (hand Hand) priority2() string {
	stackedCards := make(map[rune]int, 0)
	cardRanks := make([]rune, 0, len(hand.Cards))

	jokers := 0
	for _, card := range hand.Cards {
		cardRanks = append(cardRanks, CamelCardStrength2[card])
		if card == 'J' {
			jokers++
			continue
		}
		if _, exists := stackedCards[card]; !exists {
			stackedCards[card] = 1
		} else {
			stackedCards[card]++
		}
	}

	// add jokers to largest stack
	highestCount := 0
	highestCard := hand.Cards[0]
	for card, count := range stackedCards {
		if count > highestCount {
			highestCount = count
			highestCard = card
		}
	}
	stackedCards[highestCard] += jokers
	highestCount += jokers

	multiplier := 'a'

	if highestCount == 5 {
		multiplier = 'g' // 5 of a kind
	} else if highestCount == 4 {
		multiplier = 'f' // 4 of a kind
	} else if highestCount == 3 {
		if len(stackedCards) == 2 {
			multiplier = 'e' // full house
		} else {
			multiplier = 'd' // 3 of a kind
		}
	} else if highestCount == 2 {
		if len(stackedCards) == 3 {
			multiplier = 'c' // 2 pair
		} else {
			multiplier = 'b' // 1 pair
		}
	} else {
		multiplier = 'a' // high card
	}

	return string(multiplier) + string(cardRanks)
}

func byPriority(a, b interface{}) int {
	priorityA := a.(Hand).priority()
	priorityB := b.(Hand).priority()
	return -utils.StringComparator(priorityA, priorityB)
}

func byPriority2(a, b interface{}) int {
	priorityA := a.(Hand).priority2()
	priorityB := b.(Hand).priority2()
	return -utils.StringComparator(priorityA, priorityB)
}

func day7part1() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	hands := pq.NewWith(byPriority)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(parts[1])

		hands.Enqueue(Hand{
			Cards: []rune(parts[0]),
			Bid:   bid,
		})
	}

	result := 0
	handCount := hands.Size()

	rank := 0
	for hand, ok := hands.Dequeue(); ok; hand, ok = hands.Dequeue() {
		handStruct := hand.(Hand)
		realRank := handCount - rank
		// fmt.Println("rank", realRank, "hand", string(handStruct.Cards), "bid", handStruct.Bid, "sort", handStruct.priority())
		result += realRank * handStruct.Bid
		rank++
	}

	fmt.Println("result:", result)
}

func day7part2() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	hands := pq.NewWith(byPriority2)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(parts[1])

		hands.Enqueue(Hand{
			Cards: []rune(parts[0]),
			Bid:   bid,
		})
	}

	result := 0
	handCount := hands.Size()

	rank := 0
	for hand, ok := hands.Dequeue(); ok; hand, ok = hands.Dequeue() {
		handStruct := hand.(Hand)
		realRank := handCount - rank
		// fmt.Println("rank", realRank, "hand", string(handStruct.Cards), "bid", handStruct.Bid, "sort", handStruct.priority2())
		result += realRank * handStruct.Bid
		rank++
	}

	fmt.Println("result:", result)
}
