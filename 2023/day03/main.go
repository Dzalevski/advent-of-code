package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Coordinate struct {
	Line int
	Char int
}

var specialChars map[Coordinate]bool

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

	specialChars = make(map[Coordinate]bool)

	isSpecialChar(input)

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func isSpecialChar(input []string) {
	for i, l := range input {
		for j, c := range l {
			if string(c) == "." {
				continue
			} else if string(c) == "@" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "#" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "$" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "%" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "-" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "&" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "*" {
				specialChars[Coordinate{Line: i, Char: j}] = true
			} else if string(c) == "+" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "/" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			} else if string(c) == "=" {
				specialChars[Coordinate{Line: i, Char: j}] = false
			}
		}
	}
}

func checkSpecialChar(start, end, lineNum, numOfLines, lineLength int) bool {
	if start < 0 {
		start = 0
	}
	if end > lineLength {
		end = lineLength
	}

	// line above
	if lineNum != 0 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[Coordinate{lineNum - 1, i}]
			if ok {
				return ok
			}
		}
	}

	// in the same line
	_, ok := specialChars[Coordinate{lineNum, start}]

	if ok {
		return ok
	}

	_, ok = specialChars[Coordinate{lineNum, end}]
	if ok {
		return ok
	}

	// below the line
	if lineNum != numOfLines-1 {
		for i := start; i <= end; i++ {
			_, ok := specialChars[Coordinate{lineNum + 1, i}]
			if ok {
				return ok
			}
		}
	}
	return false
}

func partOne(lines []string) int {
	arrLength := len(lines)
	var sum int

	for k, line := range lines {
		lineLen := len(line)
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllStringIndex(line, -1)
		for _, num := range nums {
			if checkSpecialChar(num[0]-1, num[1], k, arrLength, lineLen) {
				v, err := strconv.Atoi(line[num[0]:num[1]])
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				sum += v
			}
		}
	}

	return sum
}

func partTwo(lines []string) int {
	var sum int
	for p, v := range specialChars {
		vals := []int{}
		if v {
			re := regexp.MustCompile(`\d+`)
			if p.Line >= 1 {
				numsAbove := re.FindAllStringIndex(lines[p.Line-1], -1)
				for _, num := range numsAbove {
					if (p.Char-num[0] >= -1 && p.Char-num[0] <= 1) || (p.Char-num[1] >= -2 && p.Char-num[1] <= 0) {
						val, err := strconv.Atoi(lines[p.Line-1][num[0]:num[1]])
						if err != nil {
							fmt.Println(err.Error())
						}
						vals = append(vals, val)
					}
				}
			}

			numsLine := re.FindAllStringIndex(lines[p.Line], -1)
			for _, num := range numsLine {
				if p.Char-num[1] == 0 || p.Char-num[0] == -1 {
					val, err := strconv.Atoi(lines[p.Line][num[0]:num[1]])
					if err != nil {
						fmt.Println(err.Error())
					}
					vals = append(vals, val)
				}
			}

			if len(lines) > p.Line+1 {
				numsBelow := re.FindAllStringIndex(lines[p.Line+1], -1)

				for _, num := range numsBelow {
					if (p.Char-num[0] >= -1 && p.Char-num[0] <= 1) || (p.Char-num[1] >= -2 && p.Char-num[1] <= 0) {
						val, err := strconv.Atoi(lines[p.Line+1][num[0]:num[1]])
						if err != nil {
							fmt.Println(err.Error())
						}
						vals = append(vals, val)
					}
				}
			}
		}

		if len(vals) == 2 {
			sum += vals[0] * vals[1]
		}
	}
	return sum
}
