package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// TODO: finish part 2
	input := read()
	fmt.Printf("cords: %v\n", input)
	c := parse(input)
	slices.SortFunc(c, func(c1, c2 cords) int {
		return c1.y - c2.y
	})
	fmt.Printf("cords: %v\n", c)
	ma := 0
	var mc1, mc2 cords
	for _, c1 := range c {
		for _, c2 := range c {
			a := area(c1, c2)
			fmt.Printf("c1:%+v \tc2:%+v\tarea:%d\n", c1, c2, a)
			if a > ma {
				ma = a
				mc1 = c1
				mc2 = c2
			}
		}
	}
	fmt.Printf("mc1: %v\n", mc1)
	fmt.Printf("mc2: %v\n", mc2)
	fmt.Printf("ma: %v\n", ma)

	print(c)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

type cords struct {
	x, y int
}

func area(c1, c2 cords) int {
	// fmt.Printf("c1:%+v \tc2:%+v\t\n", c1, c2)
	// fmt.Printf("c1.x:%v\tc2.x:%v\t%d\n", c1.x, c2.x, abs(c1.x-c2.x)+1)
	// fmt.Printf("c1.y:%v\tc2.y:%v\t%d\n", c1.y, c2.y, abs(c1.y-c2.y)+1)

	ret := (abs(c1.x-c2.x) + 1) * (abs(c1.y-c2.y) + 1)

	return ret
}

func abs(br int) int {
	if br < 0 {
		return -br
	}

	return br
}

func parse(input []string) []cords {
	ret := make([]cords, 0, len(input))

	for _, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		ret = append(ret, cords{x, y})
	}

	return ret
}

func print(arr []cords) {
	for y := range 9 {
		for x := range 13 {
			if slices.ContainsFunc(arr, func(c cords) bool {
				return x == c.x && y == c.y
			}) {
				fmt.Print("#")
			} else {

				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func makeField(arr []cords, size int) []cords {
	ret := make([]cords, size)

	i := 0
	for y := range size {
		// line has red tiles
		if arr[i].y == y {
			if y-1 > 0 {
				ret[y] = arr[i]
				continue
			}
		} else {
			// copy from before
			if y-1 > 0 {
				continue
			}

			ret[y] = ret[y-1]
		}
	}

	return ret
}
