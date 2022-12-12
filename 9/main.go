package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	lines := readLines()

	head := &Point{0, 0}
	tails := make([]*Point, 0)

	for i := 0; i < 9; i++ {
		t := &Point{0, 0}
		tails = append(tails, t)
	}

	total := 1
	visited := make(map[Point]bool)
	visited[*tails[len(tails)-1]] = true

	for _, line := range lines {
		direction, count := parseLine(line)
		lastTail := head
		for i := 0; i < count; i++ {
			head.move(direction)
			for _, tail := range tails {
				tail.follow(*lastTail)
				lastTail = tail
			}
			lastTail = head

			tail := tails[len(tails)-1]
			if !visited[*tail] {
				total++
				visited[*tail] = true
			}
		}
	}
	fmt.Println(total)
}

func (p *Point) follow(ref Point) {
	// se tail e head estao em contato
	if math.Abs(float64(p.y-ref.y)) <= 1 && math.Abs(float64(p.x-ref.x)) <= 1 {
		return
	}

	if p.x < ref.x {
		if p.y < ref.y {
			p.x++
			p.y++
		} else if p.y > ref.y {
			p.x++
			p.y--
		} else {
			p.x++
		}
	} else if p.x > ref.x {
		if p.y > ref.y {
			p.x--
			p.y--
		} else if p.y < ref.y {
			p.x--
			p.y++
		} else {
			p.x--
		}
	} else if p.x == ref.x {
		if p.y < ref.y {
			p.y++
		} else {
			p.y--
		}
	}
}

func (p *Point) move(d string) {
	switch d {
	case "U":
		p.y++
	case "D":
		p.y--
	case "R":
		p.x++
	case "L":
		p.x--
	}
}

func parseLine(line string) (string, int) {
	parsed := strings.Split(line, " ")
	count, _ := strconv.Atoi(parsed[1])
	return parsed[0], count
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
