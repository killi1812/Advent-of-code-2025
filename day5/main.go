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

	fmt.Printf("%+v\n", input)

	spaceIndex := slices.Index(input, "")

	ranges := input[:spaceIndex]
	foodIds := input[spaceIndex+1:]

	fmt.Printf("Id ranges:%+v\n", ranges)
	fmt.Printf("Foor Ids: %+v\n", foodIds)
	count := countFresh(foodIds, ranges)
	fmt.Printf("count: %v\n", count)

	rngs := parseRanges(ranges)
	cnt := rngs.count()
	fmt.Printf("cnt: %v\n", cnt)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func countFresh(foods, ranges []string) int {
	count := make(map[int]bool)

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		low, _ := strconv.Atoi(parts[0])
		hight, _ := strconv.Atoi(parts[1])

		for _, food := range foods {
			foodId, _ := strconv.Atoi(food)
			if foodId >= low && foodId <= hight {
				fmt.Printf("Fresh food: %v\n", foodId)
				count[foodId] = true
			}
		}
	}

	return len(count)
}

type (
	rngs []rng

	rng struct {
		low  int
		high int
	}
)

func parseRanges(ranges []string) rngs {
	ret := make([]rng, 0, len(ranges))

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		low, _ := strconv.Atoi(parts[0])
		hight, _ := strconv.Atoi(parts[1])
		ret = append(ret, rng{low: low, high: hight})
	}

	slices.SortFunc(ret, func(a, b rng) int {
		return a.low - b.low
	})

	return ret
}

func (arr rngs) count() uint64 {
	var cnt uint64 = 0

	activeRng := rng{}
	for _, r := range arr {
		fmt.Printf("rng: %+v\n", r)

		if activeRng.low == 0 && activeRng.high == 0 {
			activeRng = r
		} else if activeRng.high >= r.low {
			if activeRng.high <= r.high {
				fmt.Println("New High")
				activeRng.high = r.high
				fmt.Printf("active: %+v\n", activeRng)
			}
		} else {
			fmt.Println("Calculating")
			diff := activeRng.high - activeRng.low + 1
			fmt.Printf("%d - %d = %v\n", activeRng.high, activeRng.low, diff)
			cnt += uint64(diff)
			activeRng = r
		}

	}

	fmt.Println("Calculating")
	diff := activeRng.high - activeRng.low + 1
	fmt.Printf("%d - %d = %v\n", activeRng.high, activeRng.low, diff)
	cnt += uint64(diff)

	return cnt
}
