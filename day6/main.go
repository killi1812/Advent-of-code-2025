package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := read()

	nums, ops := parseV2(input)
	print(nums, ops)

	rez := calculate(nums, ops)
	fmt.Printf("Result: %v\n", rez)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parse(input []string) ([][]uint, []string) {
	rowCnt := len(input)
	numbers := make([][]uint, rowCnt-1)
	operations := make([]string, 0)

	for i, line := range input {
		parts := strings.Fields(line)
		fmt.Printf("Working on line: \t %v\n", parts)
		for _, part := range parts {
			fmt.Printf("part: %v\n", part)
			tmp, err := strconv.ParseUint(part, 10, 32)
			if err != nil {
				operations = append(operations, part)
				continue
			}
			num := uint(tmp)
			numbers[i] = append(numbers[i], num)
		}
	}

	return numbers, operations
}

func parseV2(input []string) ([][]uint, []string) {
	rowCnt := len(input)
	columnCnt := len(input[0])
	fmt.Printf("size %dx%d\n", columnCnt, rowCnt)

	numbers := make([][]uint, columnCnt)
	operations := make([]string, 0)

	i := 0
	b := strings.Builder{}
	var op byte = ' '
	for clmn := columnCnt - 1; clmn >= 0; clmn-- {
		b.Reset()

		for row := range rowCnt {
			if len(input[row]) <= clmn {
				fmt.Printf("len: %d \t clmd: %d\n", len(input[row]), clmn)
				continue
			}

			if input[row][clmn] == ' ' {
				continue
			}

			if input[row][clmn] == '*' || input[row][clmn] == '+' {
				op = input[row][clmn]
			} else {
				b.WriteByte(input[row][clmn])
			}
		}

		num, err := strconv.ParseUint(b.String(), 10, 32)
		if err != nil {
			continue
		}

		fmt.Printf("Num: %v\n", num)
		numbers[i] = append(numbers[i], uint(num))

		if op != ' ' {
			fmt.Printf("op: %v\n", op)
			operations = append(operations, string(op))
			fmt.Printf("i: %v\n", i)
			i++
			op = ' '
		}

	}

	return numbers, operations
}

func print(nums [][]uint, ops []string) {
	for i, line := range nums {
		fmt.Print(i, ".")
		for _, num := range line {
			fmt.Printf("\t%d", num)
		}
		fmt.Println()
	}
	fmt.Printf("Ops: ")
	for _, op := range ops {
		fmt.Printf("\t%s", op)
	}
	fmt.Println()
}

func calculate(nums [][]uint, ops []string) uint {
	var sum uint = 0
	rezs := make([]uint, len(ops))

	for i, op := range ops {
		for _, num := range nums[i] {
			fmt.Printf("Working on line: \t %v\n", num)
			switch op {
			case "*":
				if rezs[i] == 0 {
					rezs[i] = 1
				}

				rezs[i] *= num
			case "+":
				rezs[i] += num
			}
		}
		fmt.Printf("Result of op: %s :\t %d\n", op, rezs[i])
	}

	for _, num := range rezs {
		sum += num
	}

	return sum
}
