package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readLines()

	best := 0
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[0])-1; x++ {
			target, _ := strconv.Atoi(string(lines[y][x]))
			score := getScore(x, y, target, lines)
			if score > best {
				best = score
			}
		}
	}

	fmt.Println(best)
}

func getScore(x, y, target int, lines []string) int {
	total := 1

	curScore := 0
	for i := y + 1; i < len(lines); i++ {
		cur, _ := strconv.Atoi(string(lines[i][x]))
		curScore++
		if cur >= target {
			break
		}
	}
	total *= curScore

	curScore = 0
	for i := y - 1; i >= 0; i-- {
		cur, _ := strconv.Atoi(string(lines[i][x]))
		curScore++
		if cur >= target {
			break
		}
	}
	total *= curScore

	curScore = 0
	for j := x + 1; j < len(lines[0]); j++ {
		cur, _ := strconv.Atoi(string(lines[y][j]))
		curScore++
		if cur >= target {
			break
		}
	}
	total *= curScore

	curScore = 0
	for j := x - 1; j >= 0; j-- {
		cur, _ := strconv.Atoi(string(lines[y][j]))
		curScore++
		if cur >= target {
			break
		}
	}
	total *= curScore

	return total
}

func readLines() []string {
	f, _ := os.Open("input.txt")
	//f, _ := os.Open("input_test.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
