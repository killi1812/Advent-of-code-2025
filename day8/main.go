package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := read()
	cords := parse(input)

	mcir.Connect(Closest(cords))
	fmt.Printf("mcir: %+v\n", mcir)

	mcir.Connect(Closest(cords))
	fmt.Printf("mcir: %+v\n", mcir)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

type poz struct {
	x, y, z     int
	connections []*poz
}

type circuit struct {
	connections []poz
}

var mcir circuit = circuit{}

func (c circuit) IsConnected(key, value poz) bool {
	return slices.IndexFunc(c.connections, func(p2 circuit) bool {
		return key.Distance(p2.value) == 0.0
	}) != -1
}

func (c circuit) Connect(key, value poz) {
	c.connections
}

func (p poz) Distance(p2 poz) float64 {
	rez := 0.0
	rez = math.Sqrt(math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2) + math.Pow(float64(p2.z-p.z), 2))
	return rez
}

func Closest(input []poz) (poz, poz) {
	pair := make([]poz, 2)
	arr := make([]poz, len(input))

	var min poz = arr[0]
	var index int
	for i, p := range arr {

		fmt.Printf("Working with \tp:\t %v\t", p)
		ind := slices.IndexFunc(arr, func(p2 poz) bool {
			return p.Distance(p2) == 0.0
		})
		clean := exclude(arr, ind)

		tmp := slices.MinFunc(clean, func(a, b poz) int {
			return int(p.Distance(a)) - int(p.Distance(b))
		})

		fmt.Printf("Closest point:\t %v\t %+v\n", tmp.Distance(p), tmp)

		if p.Distance(tmp) < p.Distance(min) && !mcir.IsConnected(p, tmp) {
			min = tmp
			index = i
		}

		copy(arr, input)
	}

	pair[0] = arr[index]

	p2i := slices.IndexFunc(arr, func(p2 poz) bool {
		return min.Distance(p2) == 0.0
	})
	pair[1] = arr[p2i]

	return pair[0], pair[1]
}

func NewPoz(x, y, z int) poz {
	return poz{
		x: x,
		y: y,
		z: z,
	}
}

func parse(input []string) []poz {
	arr := make([]poz, len(input))

	for i, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		arr[i] = NewPoz(x, y, z)
	}

	return arr
}

func exclude(arr []poz, index int) []poz {

	if len(arr)-1 == index {
		return arr[:index-1]
	}
	return append(arr[:index], arr[index+1:]...)
}
