package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main1() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

  names := map[string]string{
    "X": "rock",
    "Y": "paper",
    "Z": "scissors",
    "A": "rock",
    "B": "paper",
    "C": "scissors",
  }

	score := 0
	for _, line := range lines {
		parsed := strings.Split(line, " ")

		a := parsed[0]
		b := parsed[1]

		cur := 0
		if names[a] == names[b] {
			cur += 3
		} else {
			if a == "A" && b == "Y" {
				cur += 6
			} else if a == "B" && b == "Z" {
				cur += 6
			} else if a == "C" && b == "X" {
        cur += 6
			}
		}

		cur += points[b]
		score += cur
	}

  fmt.Println(score)
}
