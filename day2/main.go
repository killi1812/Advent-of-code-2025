package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int

func main() {
	input := read()
	cnt := len(input)
	sums := make(chan int64, cnt)

	for _, r := range input {
		low, high := parse(r)
		go sum(low, high, sums)
	}

	var sum int64
	doneCnt := 0
	for {
		sum += <-sums

		doneCnt++
		fmt.Printf("[%d/%d]\n", doneCnt, cnt)

		if doneCnt == cnt {
			close(sums)
			break
		}
	}

	fmt.Printf("Result: %d\n", sum)
	fmt.Printf("Count: %v\n", count)
	fmt.Printf("Off raget: %d\n", 11323661261-sum)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), ",")
}

func sum(low, high int64, rez chan int64) {
	// fmt.Printf("Working on range %d-%d\n", low, high)
	var sum int64
	for i := low; i <= high; i++ {
		br := strconv.FormatInt(i, 10)

		// if p1Logic(br) {
		// 	continue
		// }

		if p2Logic(br) {
			continue
		}

		// fmt.Printf("Found Invalid ID: %d\n", i)
		sum += i
	}

	rez <- sum
	// fmt.Printf("Done with range %d-%d\n", low, high)
}

func p1Logic(br string) bool {
	if len(br)%2 != 0 {
		return true
	}

	hindex := len(br) / 2

	return br[:hindex] != br[hindex:]
}

func parse(s string) (int64, int64) {
	rez := strings.Split(s, "-")
	if len(rez) != 2 {
		fmt.Printf("Error can't split: %s\n", s)
		return 0, 0
	}
	br1, err := strconv.ParseInt(rez[0], 10, 64)
	if err != nil {
		fmt.Printf("Error convert to int: %s\n", rez[0])
		return 0, 0
	}

	br2, err := strconv.ParseInt(rez[1], 10, 64)
	if err != nil {
		fmt.Printf("Error convert to int: %s\n", rez[1])
		return 0, 0
	}

	return br1, br2
}

// p2Logic return true if br is valid id
func p2Logic(br string) bool {
	splitStr := ""

OUTER_LOOP:
	for char := range br {

		splitStr += string(br[char])

		parts := strings.Split(br, splitStr)

		// remove empty entry
		parts = parts[1:]
		// add Split prefix
		for i := range parts {
			parts[i] = splitStr + parts[i]
		}

		if len(parts) < 2 {
			continue
		}

		for _, part := range parts[1:] {
			if part != parts[0] {
				continue OUTER_LOOP
			}
		}

		fmt.Printf("Number: %v \t Split Str: %v\t Parts: %+v\n", br, splitStr, parts)
		count++
		return false
	}

	return true
}
