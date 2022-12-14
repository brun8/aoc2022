package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation func(int, int) int
	divisible int
}

func main() {
	lines := readLines()

	start := 0
	parsed := make([]*Monkey, 0)
	for i, line := range lines {
		if len(line) == 0 {
			parsed = append(parsed, parseMonkey(lines[start:i]))
			fmt.Println(lines[start:i])
			start = i + 1
		}
	}
	parsed = append(parsed, parseMonkey(lines[start:len(lines)-1]))

	for _, m := range parsed {
		fmt.Printf("%+v\n", m)
	}
}

func parseMonkey(lines []string) *Monkey {
	m := &Monkey{}

	items := parseItems(lines[1])
	m.items = items

	return m
}

func parseOperation(line string) func(int) int {
	exp := strings.Split(line, " = ")[1]
	tokens := strings.Split(exp, " ")
	return nil
}

func parseItems(line string) []int {
	items := make([]int, 0)
	nums := strings.Split(line, ": ")[1]
	for _, s := range strings.Split(nums, ", ") {
		n, _ := strconv.Atoi(s)
		items = append(items, n)
	}
	return items
}

func readLines() []string {
	//f, _ := os.Open("input.txt")
	f, _ := os.Open("input_test.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
