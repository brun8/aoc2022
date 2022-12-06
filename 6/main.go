package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines()

	for _, line := range lines {
		for i := 0; i < len(line)-13; i++ {
			m := make(map[byte]bool)
			t := 0
			for j := i; j < i+14; j++ {
				if !m[line[j]] {
					t++
					m[line[j]] = true
				}
			}
			if t == 14 {
				fmt.Println(i + 14)
				break
			}
		}
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
