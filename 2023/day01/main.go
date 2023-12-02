package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func partOne(input []string) int {
	res := 0

	for _, line := range input {
		firstDigit, lastDigit := 0, 0

		// Find the first digit in the string starting from left
		for _, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				firstDigit = digit
				break
			}
		}

		// Find the last digit in the string (starting from the right)
		for i := len(line) - 1; i >= 0; i-- {
			if digit, err := strconv.Atoi(string(line[i])); err == nil {
				lastDigit = digit
				break
			}
		}

		concatenatedValue := firstDigit*10 + lastDigit
		res += concatenatedValue
	}

	return res
}

func partTwo(input []string) int {
	totalSum := 0

	for _, line := range input {
		fd, ld := 0, 0

		fd = findFirstDigit(line)
		ld = findLastDigit(line)

		concatValue := fd*10 + ld
		totalSum += concatValue
	}
	return totalSum
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read file input %w", err))
		return
	}
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func findLastDigit(s string) int {
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Initialize variables to store the last found digit and its position
	var lastFoundDigit int
	var lastFoundPosition int

	// Check string digits
	for i, digit := range stringDigits {
		index := strings.LastIndex(s, digit)
		if index > lastFoundPosition {
			// Found a string digit with a higher position
			lastFoundPosition = index
			lastFoundDigit = i + 1
		}
	}

	// Check number digits
	for i, char := range s {
		if unicode.IsDigit(char) && i > lastFoundPosition {
			// Found a number digit with a higher position
			lastFoundPosition = i
			lastFoundDigit, _ = strconv.Atoi(string(char))
		}
	}

	if lastFoundPosition == -1 {
		log.Println("No digit found")
		return 0
	}

	return lastFoundDigit
}

func findFirstDigit(s string) int {
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Initialize variables to store the first found digit and its position
	var firstFoundDigit int
	var firstFoundPosition int

	// Check string digits
	for i, digit := range stringDigits {
		index := strings.Index(s, digit)
		if index != -1 && (firstFoundPosition == 0 || index < firstFoundPosition) {
			// Found a string digit with a lower position or it's the first one found
			firstFoundPosition = index
			firstFoundDigit = i + 1
		}
	}

	// Check number digits
	for i, char := range s {
		if unicode.IsDigit(char) && (firstFoundPosition == 0 || i < firstFoundPosition) {
			// Found a number digit with a lower position or it's the first one found
			firstFoundPosition = i
			firstFoundDigit, _ = strconv.Atoi(string(char))
		}
	}

	if firstFoundPosition == len(s) {
		log.Println("No digit found")
		return 0
	}

	return firstFoundDigit
}
