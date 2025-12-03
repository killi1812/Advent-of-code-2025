package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := read()
	var sum uint64

	for _, bank := range input {
		sum += FindBiggest(bank)
	}
	fmt.Printf("Result: %d\n", sum)
	fmt.Printf("Offset: %d\n", sum-3121910778618)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func FindBiggest(str string) uint64 {
	arr := []rune(str)
	if len(arr) < 12 {
		return 0
	}

	strNum := strings.Builder{}
	strNum.Grow(12)
	digetIndex := 0
	fmt.Printf("Bank: %v \n ", str)
	start := 0

	for strNum.Len() != 12 {

		start += digetIndex
		end := len(arr) - 11 + strNum.Len()
		if end > len(arr) {
			end = len(arr) - 1
		}

		sli := arr[start:end]
		fmt.Printf("Working on slice: %v\n", string(sli))
		diget := slices.Max(sli)
		digetIndex = slices.Index(sli, diget)
		if digetIndex == -1 {
			fmt.Printf("Error finxing number %v in %v\n", diget, sli)
			return 0
		}

		fmt.Printf("diget: %v on Index: %v\n", string(diget), digetIndex)
		digetIndex++

		strNum.WriteRune(diget)
	}

	num, err := strconv.ParseUint(strNum.String(), 10, 64)
	if err != nil {
		fmt.Printf("Error converting num %v\n", strNum)
		return 0
	}

	fmt.Printf("Bank: %v \t Max: %v\n", str, num)
	return num
}
