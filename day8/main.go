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

	for range 10 {
		p1, p2 := Closest(cords)
		p1.Connect(p2)
		fmt.Printf("p1: %+v\n", p1)
		fmt.Printf("p2: %+v\n", p2)
		fmt.Printf("%v\n", p1.IsConnected(*p2))
		fmt.Printf("%v\n", p2.IsConnected(*p1))
	}
	fmt.Printf("cords: %v\n", cords)
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

func (p poz) IsConnected(value poz) bool {
	rez := slices.IndexFunc(p.connections, func(p2 *poz) bool {
		return value.Distance(*p2) == 0.0
	}) != -1

	fmt.Printf("\nChecking connection between %+v and %+v\t%v\n", p, value, rez)
	return rez
}

func (p *poz) Connect(p2 *poz) {
	if p.IsConnected(*p2) {
		return
	}

	p.connections = append(p.connections, p2)
	p2.connections = append(p2.connections, p)
}

func (p poz) Distance(p2 poz) float64 {
	rez := 0.0
	rez = math.Sqrt(math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2) + math.Pow(float64(p2.z-p.z), 2))
	return rez
}

func Closest(input []poz) (*poz, *poz) {
	min := math.MaxFloat32
	var index1 int
	var index2 int

	for i, p := range input {

		fmt.Printf("Working with \t\tp:\t %v\t\n", p)
		fmt.Println("----------------------Calling min function----------------------")
		tmp := slices.MinFunc(exclude(input, i), func(a, b poz) int {
			fmt.Printf("a: %v\n", a)
			fmt.Printf("b: %v\n", b)
			if p.IsConnected(a) {
				return int(p.Distance(b))
			}

			if p.IsConnected(b) {
				return int(p.Distance(a))
			}

			return int(p.Distance(b)) - int(p.Distance(a))
		})

		fmt.Printf("Closest point:\t\t %v\t %v\n", tmp.Distance(p), tmp)

		fmt.Printf("%v Distance(%v): %v\n", p, tmp, p.Distance(tmp))

		if val := p.Distance(tmp); val < min {
			min = val
			index1 = i
			index2 = slices.IndexFunc(input, func(p2 poz) bool {
				return tmp.Distance(p2) == 0.0
			})

		}
	}

	return &input[index1], &input[index2]
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

func exclude(input []poz, index int) []poz {
	arr := make([]poz, len(input))
	copy(arr, input)

	if len(arr)-1 == index {
		fmt.Printf("rez: %v\n", arr[:index])
		return arr[:index]
	}
	rez := append(arr[:index], arr[index+1:]...)
	fmt.Printf("rez: %v\n", rez)
	return rez
}
