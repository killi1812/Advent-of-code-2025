package main

import (
	"fmt"
	"os"
	"strings"
)

type Mashine struct {
	state []bool

	correct []bool
	buttons [][]int
	joltage []int
}

func NewMashine(c, b, j string) Mashine {
	fmt.Printf("c: %v\n", c)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("j: %v\n", j)

	ret := Mashine{}
	return ret
}

func main() {
	input := read()
	parse(input)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parse(input []string) []Mashine {
	ret := make([]Mashine, len(input))

	for i, line := range input {

		i0 := strings.IndexRune(line, ' ')
		i1 := strings.IndexRune(line, '{')

		ret[i] = NewMashine(line[:i0], line[i0+1:i1], line[i1:])
	}

	return ret
}
