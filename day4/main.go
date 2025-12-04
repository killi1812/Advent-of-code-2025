package main

import (
	"fmt"
	"os"
	"strings"
)

const empty rune = '.'
const crate rune = '@'
const crate_can_move rune = 'x'

type werhouse struct {
	arr  [][]rune
	lenX int
	lenY int
}

func main() {
	input := read()
	w := newWerhouse(input)
	w.print()

	count := 0
	for {
		w = w.MarkMovable()
		tmp := w.Count()
		count += tmp
		w.print()
		w = w.RemoveMovable()
		w.print()
		if tmp == 0 {
			break
		}
	}

	fmt.Printf("Result: %v\n", count)
}

func read() [][]rune {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	ret := make([][]rune, 0, len(lines))

	for _, line := range lines {
		ret = append(ret, []rune(line))
	}

	return ret
}

func (w werhouse) print() {
	for _, line := range w.arr {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}

	fmt.Printf("\n----------------------------------------------------------------------\n\n")
}

func newWerhouse(arr [][]rune) werhouse {
	ret := werhouse{
		arr:  arr,
		lenX: len(arr),
		lenY: len(arr[0]),
	}
	return ret
}

func (w werhouse) Count() int {
	count := 0
	for _, line := range w.arr {
		for _, char := range line {
			if char == crate_can_move {
				count++
			}
		}
	}
	return count
}

func (w werhouse) MarkMovable() werhouse {
	for y, line := range w.arr {
		for x := range line {
			if w.CanMove(x, y) {
				w.arr[y][x] = crate_can_move
			}
		}
	}

	return w
}

func (w werhouse) RemoveMovable() werhouse {
	for y, line := range w.arr {
		for x := range line {
			if w.arr[y][x] == crate_can_move {
				w.arr[y][x] = empty
			}
		}
	}

	return w
}

func (w werhouse) CanMove(x, y int) bool {
	count := 0
	if w.arr[y][x] != crate {
		return false
	}

	// check for upper y bound
	if y > 0 {
		if w.arr[y-1][x] != empty {
			count++
		}

		// check upper left bound
		if x > 0 {
			if w.arr[y-1][x-1] != empty {
				count++
			}
		}

		// check upper right bound
		if x < w.lenX-1 {
			if w.arr[y-1][x+1] != empty {
				count++
			}
		}

	}

	// check right bound
	if x > 0 {
		if w.arr[y][x-1] != empty {
			count++
		}
	}

	// check right bound
	if x < w.lenX-1 {
		if w.arr[y][x+1] != empty {
			count++
		}
	}

	// check lower bound
	if y < w.lenY-1 {
		if w.arr[y+1][x] != empty {
			count++
		}

		// check lower left bound
		if x > 0 {
			if w.arr[y+1][x-1] != empty {
				count++
			}
		}

		// check lower right bound
		if x < w.lenX-1 {
			if w.arr[y+1][x+1] != empty {
				count++
			}
		}

	}

	return count < 4
}
