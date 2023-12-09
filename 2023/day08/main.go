package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	input = strings.TrimRight(input, "\n")

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

type TreeNode struct {
	key   string
	left  *TreeNode
	right *TreeNode
}

func partOne(input string) any {
	split := strings.Split(input, "\n")
	dirs := split[0]

	nodes := make(map[string]*TreeNode)

	for _, line := range split[2:] {
		ckey, lkey, rkey := line[0:3], line[7:10], line[12:15]
		left := createNodeIfAbsent(nodes, lkey, nil, nil)
		right := createNodeIfAbsent(nodes, rkey, nil, nil)
		createNodeIfAbsent(nodes, ckey, left, right)
	}

	return travelAcross(nodes["AAA"], dirs, func(k string) bool { return k == "ZZZ" })
}

func partTwo(input string) any {
	split := strings.Split(input, "\n")
	dirs := split[0]

	nodes := make(map[string]*TreeNode)

	roots := make([]*TreeNode, 0)
	// continue after the new line
	for _, line := range split[2:] {
		ckey, leftKey, rightKey := line[0:3], line[7:10], line[12:15]

		left := createNodeIfAbsent(nodes, leftKey, nil, nil)
		right := createNodeIfAbsent(nodes, rightKey, nil, nil)
		root := createNodeIfAbsent(nodes, ckey, left, right)
		if ckey[len(ckey)-1] == 'A' {
			roots = append(roots, root)
		}
	}

	res := 1
	for _, root := range roots {
		res = leastCommonMultiple(res, travelAcross(root, dirs, func(k string) bool { return k[len(k)-1] == 'Z' }))
	}

	return res
}

func createNodeIfAbsent(nodes map[string]*TreeNode, key string, left, right *TreeNode) *TreeNode {
	if nodes[key] == nil {
		nodes[key] = &TreeNode{
			key:   key,
			left:  left,
			right: right,
		}
	} else {
		if left != nil {
			nodes[key].left = left
		}
		if right != nil {
			nodes[key].right = right
		}
	}

	return nodes[key]
}

func travelAcross(root *TreeNode, dirs string, isFinish func(k string) bool) int {
	cur := root
	curI := 0
	cnt := 0

	for cur != nil && !isFinish(cur.key) {
		cnt++
		if dirs[curI] == 'L' {
			cur = cur.left
		} else {
			cur = cur.right
		}

		curI = (curI + 1) % len(dirs)
	}

	return cnt
}

func leastCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
