package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var hands []Hand
	for _, l := range input {

		s := strings.Fields(l)
		bid, _ := strconv.Atoi(s[1])

		h := Hand{
			cards:     s[0],
			bid:       bid,
			handType:  getHandType(s[0]),
			handType2: getHandTypeWithJokers(s[0]),
		}
		hands = append(hands, h)
	}

	fmt.Println(partOne(hands))
	fmt.Println(partTwo(hands))
}

type Hand struct {
	cards     string
	bid       int
	handType  handType
	handType2 handType
}

type handType int

const (
	FiveOfKind  handType = 7
	FourOfKind           = 6
	FullHouse            = 5
	ThreeOfKind          = 4
	TwoPair              = 3
	OnePair              = 2
	HighCard             = 1
)

func getHandType(cards string) handType {
	freq := map[rune]int{}
	for _, c := range cards {
		if count, ok := freq[c]; ok {
			count++
			freq[c] = count
		} else {
			freq[c] = 1
		}
	}

	for _, v := range freq {
		if v == 5 {
			return FiveOfKind
		}
		if v == 4 {
			return FourOfKind
		}
	}
	if len(freq) == 2 {
		return FullHouse
	}
	for _, v := range freq {
		if v == 3 {
			return ThreeOfKind
		}
	}
	{
		c := 0
		for _, v := range freq {
			if v == 2 {
				c++
			}
		}
		if c == 2 {
			return TwoPair
		} else if c == 1 {
			return OnePair
		}
	}
	return HighCard
}

func getHandTypeWithJokers(cards string) handType {
	freq := map[rune]int{}
	for _, c := range cards {
		if count, ok := freq[c]; ok {
			count++
			freq[c] = count
		} else {
			freq[c] = 1
		}
	}

	jokers := freq['J']
	if jokers == 0 {
		return getHandType(cards)
	}

	var mk rune
	var mv int
	for k, v := range freq {
		if k == 'J' {
			continue
		}
		if v > mv {
			mk = k
			mv = v
		}
	}

	cards = strings.ReplaceAll(cards, "J", string(mk))
	return getHandType(cards)
}

func cardTypes(a rune, b rune) int {
	types := []rune{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}
	if a == b {
		return 0
	}
	var xIdx, yIdx int
	for i, v := range types {
		if v == a {
			xIdx = i
		}
		if v == b {
			yIdx = i
		}
	}
	if xIdx < yIdx {
		return 1
	}
	return -1
}

func cardTypesWithJoker(a rune, b rune) int {
	types := []rune{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
	}
	if a == b {
		return 0
	}
	var xIdx, yIdx int
	for i, v := range types {
		if v == a {
			xIdx = i
		}
		if v == b {
			yIdx = i
		}
	}
	if xIdx < yIdx {
		return 1
	}
	return -1
}

func partOne(hands []Hand) int {
	cardsTypes := func(a, b Hand) int {
		if a.handType < b.handType {
			return -1
		}
		if a.handType > b.handType {
			return 1
		}
		for i := range a.cards {
			c := cardTypes(rune(a.cards[i]), rune(b.cards[i]))
			if c != 0 {
				return c
			}

		}
		return 0
	}
	slices.SortFunc(hands, cardsTypes)

	ans := 0
	for i, h := range hands {
		rank := i + 1
		ans += rank * h.bid
	}
	return ans
}

func partTwo(hands []Hand) int {
	cardTypes := func(a, b Hand) int {
		if a.handType2 < b.handType2 {
			return -1
		}
		if a.handType2 > b.handType2 {
			return 1
		}
		for i := range a.cards {
			c := cardTypesWithJoker(rune(a.cards[i]), rune(b.cards[i]))
			if c != 0 {
				return c
			}

		}
		return 0
	}
	slices.SortFunc(hands, cardTypes)

	ans := 0
	for i, h := range hands {
		rank := i + 1
		ans += rank * h.bid
	}
	return ans
}
