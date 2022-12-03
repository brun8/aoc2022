package main

import (
	"bufio"
	"fmt"
	"os"
)

// eu nao tenho orgulho nenhum do que foi feito aqui no dia de hoje
func main() {
  f, _ := os.Open("input.txt")
	//f, _ := os.Open("input_test.txt")
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

  total := 0
	for i := 0; i < len(lines); i += 3 {
    cur := 0
		found1 := make(map[rune]bool)
		found2 := make(map[rune]bool)
		for _, c := range lines[i] {
			found1[c] = true
		}
		for _, c := range lines[i+1] {
			found2[c] = true
		}
		for _, c := range lines[i+2] {
			if found1[c] && found2[c] {
        cur += getPriority(c)
        break
			}
		}
    total += cur
	}
  fmt.Println(total)
}

func split(s string) (string, string) {
	return s[0 : len(s)/2], s[len(s)/2:]
}

func getPriority(c rune) int {
	var priority int
	val := int(c)
	if val >= 97 {
		priority = val - 96 // 96 valor do rune 'a' - 1. prioridade('a') = 1
	} else {
		priority = val - 64 + 26 // prioridade('A') = 27
	}
	return priority
}
