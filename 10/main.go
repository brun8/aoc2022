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

	cycle := 0
	x := 1
	cdt := 0
	res := ""

	for _, line := range lines {
		cycle++
		fmt.Println(x)
		if cdt%40 == x || cdt%40 == x+1 || cdt%40 == x-1 {
			res += "#"
		} else {
			res += " "
		}
		cdt++

		parsed := strings.Split(line, " ")

		if parsed[0] == "addx" {
			val, _ := strconv.Atoi(parsed[1])
			cycle++
			if cdt%40 == x || cdt%40 == x+1 || cdt%40 == x-1 {
				res += "#"
			} else {
				res += " "
			}
			cdt++
			x += val
		}
	}

	for i := 0; i < len(res); i += 40 {
		fmt.Println(res[i : i+40])
	}
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
