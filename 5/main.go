package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines()
	y, x := separateInput(lines)
	crates := parseCrates(y)

	for _, command := range x {
		a, b, c := parseCommand(command)
		target := crates[c-1]
		removed_list := make([]rune, 0)

		for i := 0; i < a; i++ {
			source := crates[b-1]
			removed := source[len(source)-1]
			new_source := source[:len(source)-1]

			removed_list = append(removed_list, removed)
			crates[b-1] = new_source
		}

		removed_list = reverse(removed_list)
		for _, removed := range removed_list {
			target = append(target, removed)
			crates[c-1] = target
		}
	}

	s := ""
	for _, crate := range crates {
		if len(crate) > 0 {
			s += string(crate[len(crate)-1])
		}
	}
	fmt.Println(s)
}

func parseCommand(c string) (int, int, int) {
	arr := strings.Split(c, " ")
	count, _ := strconv.Atoi(arr[1])
	source, _ := strconv.Atoi(arr[3])
	target, _ := strconv.Atoi(arr[5])

	return count, source, target
}

func parseCrates(lines []string) [][]rune {
	crates := make([][]rune, 0)
	for i := 0; i < (len(lines[0])+1)/4; i++ {
		crates = append(crates, []rune{})
	}

	for _, line := range lines {
		for i, c := range line {
			if c >= 'A' && c <= 'Z' {
				crate := crates[int(i/4)]
				crates[int(i/4)] = append(crate, c)
			}
		}
	}
	for i, crate := range crates {
		crates[i] = reverse(crate)
	}

	return crates
}

func reverse(a []rune) []rune {
	output := make([]rune, len(a))

	for i, n := range a {
		output[len(a)-1-i] = n
	}

	return output
}

func separateInput(lines []string) ([]string, []string) {
	sep := 0
	for i, line := range lines {
		if len(line) == 0 {
			sep = i
		}
	}
	return lines[:sep], lines[sep+1:]
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
