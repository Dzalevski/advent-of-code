package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}

func main() {

	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

type Race struct {
	Time     int
	Distance int
}

// Extract integer values from a line
func extractValues(line string) []int {
	var values []int
	fields := strings.Fields(line)
	for _, field := range fields[1:] {
		value, _ := strconv.Atoi(field)
		values = append(values, value)
	}
	return values
}

func partOne(input string) int {

	lines := strings.Split(input, "\n")

	timeValues := extractValues(lines[0])
	distanceValues := extractValues(lines[1])

	var races []Race
	for i := range timeValues {
		races = append(races, Race{Time: timeValues[i], Distance: distanceValues[i]})
	}

	var sum = 1

	for _, race := range races {
		var wayToWin int

		for i := 0; i < race.Time; i++ {
			speed := i
			remainingTime := race.Time - i
			distance := speed * remainingTime

			if distance > race.Distance {
				wayToWin++
			}

		}
		sum = wayToWin * sum

	}

	return sum
}

func concatenateValues(line string) int {
	line = strings.Split(line, ":")[1]

	line = strings.ReplaceAll(line, " ", "")

	value, _ := strconv.Atoi(line)
	return value
}

func partTwo(input string) int {
	// Split the input by lines
	lines := strings.Split(input, "\n")

	// Extract and concatenate time and distance values
	timeValue := concatenateValues(lines[0])
	distanceValue := concatenateValues(lines[1])

	var wayToWin int
	var sum = 1
	for i := 0; i < timeValue; i++ {
		speed := i
		remainingTime := timeValue - i
		distance := speed * remainingTime

		if distance > distanceValue {
			wayToWin++
		}

	}
	sum = wayToWin * sum

	return sum
}
