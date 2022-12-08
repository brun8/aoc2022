package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	draw := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	loose := map[string]int{
		"A": 3,
		"B": 1,
		"C": 2,
	}

	win := map[string]int{
		"A": 2,
		"B": 3,
		"C": 1,
	}

	score := 0
	for _, line := range lines {
		parsed := strings.Split(line, " ")

		a := parsed[0]
		b := parsed[1]

		cur := 0

		if b == "X" {
			cur += loose[a]
		} else if b == "Y" {
			cur += draw[a] + 3
		} else {
			cur += win[a] + 6
		}
		score += cur
	}

	fmt.Println(score)
}
