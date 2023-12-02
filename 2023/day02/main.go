package main

import (
	"bufio"
	"fmt"
	"os"
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

func partOne(input []string) int {
	var (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
		total    = 0
	)

	possibleGames := make(map[int]bool)

	for _, line := range input {
		game := strings.Split(line, ":")
		gameID, err := strconv.Atoi(strings.TrimPrefix(game[0], "Game "))
		if err != nil {
			fmt.Println(fmt.Errorf("error while parsing line %s: %w", game[0], err))
			return 0
		}

		possibleGames[gameID] = true

		sets := strings.Split(game[1], ";")

		for _, set := range sets {
			colorsInSets := strings.Split(set, ",")
			colorsCount := make(map[string]int)

			for _, colorsInSet := range colorsInSets {
				parts := strings.Fields(colorsInSet)
				colorName := parts[1]
				colorNumberStr := parts[0]

				colorNumber, err := strconv.Atoi(colorNumberStr)
				if err != nil {
					fmt.Println(fmt.Errorf("failed to parse color number string: %w", err))
					return 0
				}

				colorsCount[colorName] = colorNumber
			}

			if colorsCount["red"] > maxRed || colorsCount["green"] > maxGreen || colorsCount["blue"] > maxBlue {
				possibleGames[gameID] = false
				break
			}
		}
	}

	for gameID, isPossible := range possibleGames {
		if isPossible {
			total += gameID
		}
	}

	return total

}

func partTwo(input []string) int {
	totalProduct := 0

	for _, line := range input {
		game := strings.Split(line, ":")
		sets := strings.Split(game[1], ";")
		colorCounts := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, set := range sets {
			colorsSets := strings.Split(set, ",")
			for _, colorsSet := range colorsSets {
				parts := strings.Fields(colorsSet)
				color := parts[1]
				colorNumberStr := parts[0]

				colorNumber, err := strconv.Atoi(colorNumberStr)
				if err != nil {
					fmt.Println(fmt.Errorf("failed to parse color number string: %w", err))
					return 0
				}
				if colorNumber > colorCounts[color] {
					colorCounts[color] = colorNumber
				}
			}
		}

		totalProduct += colorCounts["red"] * colorCounts["green"] * colorCounts["blue"]
	}

	return totalProduct
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
