package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	snacks := make([]int, 0)
	cur := 0

	for _, line := range lines {
		if len(line) == 0 {
			snacks = append(snacks, cur)
			cur = 0
		} else {
			num, _ := strconv.Atoi(line)
			cur += num
		}
	}

	sort.Ints(snacks)
	s := 0
	for i := 0; i < 3; i++ {
		s += snacks[len(snacks)-(i+1)]
	}
	fmt.Println(s)
}
