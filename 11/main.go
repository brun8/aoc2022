package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation string
	divisible int
	ifTrue    int
	ifFalse   int
	inspected int
}

func main() {
	lines := readLines()

	start := 0
	prod := 1

	ms := make([]*Monkey, 0)
	for i, line := range lines {
		if len(line) == 0 {
			m := parseMonkey(lines[start:i])
			prod *= m.divisible
			ms = append(ms, m)
			start = i + 1
		}
	}

	m := parseMonkey(lines[start:])
	prod *= m.divisible
	ms = append(ms, m)

	for i := 0; i < 10000; i++ {
		for _, m := range ms {
			// ids dos macacos que vao receber os items, pra eu nao modificar o
			// array enquanto eu itero por ele
			targets := make([]int, len(m.items))

			for i, item := range m.items {
				m.inspected++
				val := m.Inspect(item, prod)
				m.items[i] = val
				test := testDivisible(val, m.divisible)
				var targetId int
				if test {
					targetId = m.ifTrue
				} else {
					targetId = m.ifFalse
				}
				targets[i] = targetId
			}

			// joga items para os devidos macacos
			for i, item := range m.items {
				targetMonkey := ms[targets[i]]
				targetMonkey.items = append(targetMonkey.items, item)
			}
			m.items = []int{}
		}
	}
	sortMonkeys(ms)
	fmt.Println(ms[0].inspected * ms[1].inspected)
}

func (m Monkey) Inspect(num, prod int) int {
	n := runOperation(m.operation, num)

	return n % prod
}

func runOperation(op string, old int) int {
	parsed := strings.Split(op, " ")

	var a, b int
	if parsed[0] == "old" {
		a = old
	} else {
		n, _ := strconv.Atoi(parsed[0])
		a = n
	}

	if parsed[2] == "old" {
		b = old
	} else {
		n, _ := strconv.Atoi(parsed[2])
		b = n
	}

	if parsed[1] == "+" {
		return a + b
	} else {
		return a * b
	}
}

func sortMonkeys(ms []*Monkey) {
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].inspected > ms[j].inspected
	})
}

func printMonkeys(ms []*Monkey) {
	for i, m := range ms {
		fmt.Printf("%d: %d\n", i, m.inspected)
	}
}

func parseMonkey(lines []string) *Monkey {
	m := &Monkey{}
	m.items = parseItems(lines[1])
	m.operation = parseOperation(lines[2])
	m.divisible = getIntAtEnd(lines[3])
	m.ifTrue = getIntAtEnd(lines[4])
	m.ifFalse = getIntAtEnd(lines[5])

	return m
}

func testDivisible(a, b int) bool {
	return a%b == 0
}

func getIntAtEnd(line string) int {
	parsed := strings.Split(line, " ")
	ns := parsed[len(parsed)-1]

	n, _ := strconv.Atoi(ns)

	return n
}

func parseOperation(line string) string {
	return strings.Split(line, " = ")[1]
}

func parseItems(line string) []int {
	nums := strings.Split(line, ": ")[1]
	items := []int{}

	for _, num := range strings.Split(nums, ", ") {
		numInt, _ := strconv.Atoi(num)
		items = append(items, numInt)
	}

	return items
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
