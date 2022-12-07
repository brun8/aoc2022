package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	path     string
	size     int
	parent   *Dir
	children []*Dir
}

func main() {
	lines := readLines()
	cur := &Dir{path: "/"}
	dirs := []*Dir{cur}

	for _, line := range lines[1:] {
		arr := strings.Split(line, " ")
		if arr[0] == "$" {
			if arr[1] == "ls" {
				continue
			}

			if arr[2] == ".." {
				cur = cur.parent
				continue
			}

			newCur := &Dir{path: arr[2], parent: cur}
			cur.children = append(cur.children, newCur)
			cur = newCur
			dirs = append(dirs, cur)

			continue
		}

		if arr[1] != "dir" {
			size, _ := strconv.Atoi(arr[0])
			cur.size += size
		}
	}

	used := dirs[0].GetNestedSize()
	needed := 30000000
	free := 70000000 - used
	target := needed - free

	found := dirs[0]

	for _, dir := range dirs {
		size := dir.GetNestedSize()
		if size >= target && size < found.GetNestedSize() {
			found = dir
		}
	}
	fmt.Println(found.GetNestedSize())
}

func (d Dir) GetNestedSize() int {
	childrenSize := 0
	for _, children := range d.children {
		childrenSize += children.GetNestedSize()
	}
	return d.size + childrenSize
}

func parseLs(size, name string) int {
	isize, _ := strconv.Atoi(size)
	return isize
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
