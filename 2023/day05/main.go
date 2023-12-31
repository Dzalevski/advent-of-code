package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Range struct {
	DstStart int
	SrcStart int
	Length   int
}

type Map []Range

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	seeds, maps := parseInput(lines)

	fmt.Println(partOne(seeds, maps))
	fmt.Println(partTwo(seeds))
}

func partOne(seeds []int, maps []Map) int {
	return calculateSmallestLocation(seeds, maps)
}

func partTwo(seeds []int) int {
	seedPairs := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seedPairs = append(seedPairs, seeds[i:i+2])
	}

	length := 0
	for _, p := range seedPairs {
		length += p[1]
	}
	seeds = make([]int, length)
	j := 0
	for _, p := range seedPairs {
		for i := p[0]; i < p[0]+p[1]; i++ {
			seeds[j] = i
			j++
		}
	}

	return calculateSmallestLocation(seeds, nil)
}

func parseInput(lines []string) ([]int, []Map) {
	seeds := func() []int {
		ss := strings.Fields(lines[0][6:])
		var i []int
		for _, s := range ss {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				continue
			}

			i = append(i, n)
		}

		return i
	}()

	lines = lines[2:]

	var maps []Map
	mi := -1
	for _, line := range lines {
		if strings.Contains(line, "map:") {
			mi += 1
			maps = append(maps, Map{})
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		ds, _ := strconv.Atoi(fields[0])
		ss, _ := strconv.Atoi(fields[1])
		l, _ := strconv.Atoi(fields[2])
		maps[mi] = append(maps[mi], Range{
			DstStart: ds,
			SrcStart: ss,
			Length:   l,
		})
	}

	return seeds, maps
}

func calculateSmallestLocation(seeds []int, maps []Map) int {
	smallest := -1
	for i := range seeds {
	Maps:
		for _, m := range maps {
			for _, r := range m {
				if seeds[i] >= r.SrcStart && seeds[i] <= r.SrcStart+r.Length-1 {
					seeds[i] = seeds[i] - r.SrcStart + r.DstStart
					continue Maps
				}
			}
		}
		if smallest == -1 || seeds[i] < smallest {
			smallest = seeds[i]
		}
	}

	return smallest
}
