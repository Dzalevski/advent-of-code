package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		log.Fatal(err)
		return
	}

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(lines []string) int {
	var sum int
	for _, line := range lines {
		winningNumbers := make(map[int]bool)
		numbersOnly := strings.Split(line, ": ")
		numbers := strings.Split(numbersOnly[1], " | ")
		re := regexp.MustCompile(`\d+`)
		winningNums := re.FindAllString(numbers[0], -1)

		for _, n := range winningNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			winningNumbers[v] = true
		}

		elfNumbers := re.FindAllString(numbers[1], -1)

		var count int
		for _, n := range elfNumbers {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, ok := winningNumbers[v]

			if ok {
				count++
			}
		}
		x := pow(count)
		sum += x

	}

	return sum
}

func pow(n int) int {
	if n == 1 {
		return 1
	} else if n < 1 {
		return 0
	}

	return 2 * pow(n-1)
}

func partTwo(lines []string) int {
	processCount := make(map[int]int)

	for _, line := range lines {
		winningNumbers := make(map[int]bool)
		cards := strings.Split(line, ": ")
		numbers := strings.Split(cards[1], " | ")
		re := regexp.MustCompile(`\d+`)

		allNumbers, err := strconv.Atoi(re.FindString(cards[0]))
		if err != nil {
			fmt.Println(err.Error())
		}
		processCount[allNumbers]++

		winningNums := re.FindAllString(numbers[0], -1)

		for _, n := range winningNums {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			winningNumbers[v] = true
		}

		elfNumbers := re.FindAllString(numbers[1], -1)

		var count int

		for _, n := range elfNumbers {
			v, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, ok := winningNumbers[v]

			if ok {
				count++
				processCount[allNumbers+count] += processCount[allNumbers]
			}
		}
	}

	sum := 0
	for _, v := range processCount {
		sum += v
	}

	return sum
}
