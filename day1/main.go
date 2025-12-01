package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lock struct {
	value int
}

func (l *lock) RRot(num int) int {
	oldVal := l.value
	l.value += num
	ret := l.value / 100

	if l.value >= 100 {
		l.value %= 100
	}

	if ret != 0 {
		fmt.Printf("Extra points: %d, old value %d", ret, oldVal)
	}
	return ret
}

func (l *lock) LRot(num int) int {
	oldVal := l.value
	numstr := fmt.Sprintf("00%d", num)
	des, _ := strconv.Atoi(numstr[len(numstr)-2:])

	ret := 0
	if l.value <= des && l.value != 0 {
		ret++
	}

	l.value -= num
	ret += num / 100

	if l.value < 0 {
		tmp := (l.value % 100)
		if tmp != 0 {
			l.value = 100 + tmp
		} else {
			l.value = 0
		}
	}

	if ret != 0 {
		fmt.Printf("Extra points: %d, old value %d", ret, oldVal)
	}

	return ret
}

func main() {
	l := lock{value: 50}
	count := 0
	input := read()

	println(len(input))
	fmt.Println("\tRotation\tValue")

	for _, rot := range input {
		br, _ := strconv.Atoi(rot[1:])

		switch rot[0] {
		case 'L':
			count += l.LRot(br)
		case 'R':
			count += l.RRot(br)
		}

		if l.value == 0 {
			// count++
		}
		fmt.Printf("\t%s\t\t%d\n", rot, l.value)

	}
	println("Count of zeros: ", count)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data[:len(data)-1]), "\n")
}
