package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  lines := readLines()

  fmt.Println(lines)
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

