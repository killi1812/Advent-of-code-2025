package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	START byte = 'S'
	SPLIT byte = '^'
	EMPTY byte = '.'
	PATH  byte = '|'
)

type mainfold struct {
	arr      [][]byte
	cnt      [][]int
	xLen     int
	yLen     int
	splitCnt int

	activeLine int
}

func NewMainfold(input []string) mainfold {
	yLen := len(input)
	xLen := len(input[0])
	ret := mainfold{
		arr:        make([][]byte, yLen),
		cnt:        make([][]int, yLen),
		yLen:       yLen,
		xLen:       xLen,
		activeLine: 0,
	}

	for i, line := range input {
		ret.arr[i] = []byte(line)
		ret.cnt[i] = make([]int, xLen)
	}
	ret.cnt[0] = slices.Repeat([]int{1}, xLen)

	return ret
}

func (m *mainfold) Count() {
	for _, line := range m.cnt {
		for _, char := range line {
			fmt.Printf("%d ", char)
		}
		fmt.Println()
	}

	for _, char := range m.cnt[m.yLen-2] {
		m.splitCnt += char
	}
	fmt.Println()
}

func (m mainfold) Print() {
	for _, line := range m.arr {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}

	fmt.Println()
}

func (m *mainfold) Step() bool {
	if m.activeLine+1 == m.xLen {
		return false
	}

	// first step
	if m.activeLine == 0 {
		startIndex := slices.Index(m.arr[m.activeLine], START)
		if startIndex == -1 {
			panic("No start")
		}

		m.check(startIndex)
	} else {
		for i, char := range m.arr[m.activeLine] {
			if char == PATH {
				m.check(i)
			}
		}
	}

	m.activeLine++
	m.Print()
	return true
}

func (m *mainfold) check(position int) {
	if m.arr[m.activeLine+1][position] != SPLIT {
		m.cnt[m.activeLine+1][position] += m.cnt[m.activeLine][position]

		m.arr[m.activeLine+1][position] = PATH
		return
	}

	if position+1 < m.xLen {
		m.cnt[m.activeLine+1][position+1] = m.cnt[m.activeLine][position] +
			m.cnt[m.activeLine+1][position+1]

		m.arr[m.activeLine+1][position+1] = PATH
	}

	if position-1 > -1 {
		m.cnt[m.activeLine+1][position-1] = m.cnt[m.activeLine][position] +
			m.cnt[m.activeLine+1][position-1]

		m.arr[m.activeLine+1][position-1] = PATH
	}
}

func (m *mainfold) Go() int {
	for m.Step() {
	}
	return m.splitCnt
}

func main() {
	input := read()
	m := NewMainfold(input)
	m.Go()

	m.Count()
	fmt.Printf("Result: %v\n", m.splitCnt)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
