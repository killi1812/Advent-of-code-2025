package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// TODO: cache only works if we have a reverse algoritam
var cache map[string][]*device = make(map[string][]*device)

func main() {
	input := read()
	devs := parse(input)
	print(devs)
	allPaths := make([][]*device, 0)
	you := slices.IndexFunc(devs, func(d device) bool {
		return d.name == "svr"
		// return d.name == "you"
	})

	dfs(&devs[you], make([]*device, 0), &allPaths)

	fmt.Printf("Result: %v\n", len(allPaths))
	printAp(allPaths)

	cnt := 0
	for _, path := range allPaths {
		b1 := slices.ContainsFunc(path, func(d *device) bool {
			return d.name == "fft"
		})

		b2 := slices.ContainsFunc(path, func(d *device) bool {
			return d.name == "dac"
		})
		if b1 && b2 {
			cnt++
		}
	}
	fmt.Printf("cnt: %v\n", cnt)
	printC()
}

func printAp(allPaths [][]*device) {
	for i, path := range allPaths {
		fmt.Printf("Path %d: ", i)
		for _, dev := range path {
			fmt.Printf("%v,", dev.name)
		}
		fmt.Println()
	}
}

func printC() {
	for i, path := range cache {
		fmt.Printf("key %s: ", i)
		for _, dev := range path {
			fmt.Printf("%v,", dev.name)
		}
		fmt.Println()
	}
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

type device struct {
	name  string
	conns []*device
}

func NewDevice(name string, devs ...*device) *device {
	return &device{
		name:  name,
		conns: devs,
	}
}

func (d device) IsOut() bool {
	return d.name == "out"
}

func (d *device) Connect(dev *device) {
	d.conns = append(d.conns, dev)
}

func parse(input []string) []device {
	ret := make([]device, 0, len(input))
	devs := make(map[string]*device)

	for _, line := range input {
		parts := strings.Split(line, ":")

		dev, ok := devs[parts[0]]
		if !ok {
			devs[parts[0]] = NewDevice(parts[0])
			dev = devs[parts[0]]
		}

		for _, conn := range strings.Split(parts[1], " ") {
			if conn == " " || conn == "" {
				continue
			}

			_, ok = devs[conn]
			if !ok {
				devs[conn] = NewDevice(conn)
			}
			dev.Connect(devs[conn])
		}

	}

	for _, dev := range devs {
		ret = append(ret, *dev)
	}

	return ret
}

func dfs(src *device, path []*device, allPaths *[][]*device) {
	fmt.Printf("On device: \t %v options:", src.name)
	for _, conn := range src.conns {
		fmt.Printf("%s,", conn.name)
	}
	fmt.Println()

	if slices.ContainsFunc(path, func(d *device) bool {
		return src.name == d.name
	}) {
		fmt.Println("------------------------------Found cicle------------------------------")
		return
	}
	path = append(path, src)

	if src.IsOut() {
		*allPaths = append(*allPaths, path)

		fmt.Println("-----------------------------Found Out-----------------------------")

		for i, dev := range path[1:] {
			s := path[i:]
			printp(s)
			cache[dev.name] = s
		}

	} else {
		for _, dev := range src.conns {
			cpath, ok := cache[src.name]
			if ok && false {
				fmt.Printf("src.name: %v\n", src.name)
				fmt.Println("-----------------------------Using cache-----------------------------")
				path = append(path, cpath...)
				*allPaths = append(*allPaths, path)

			} else {
				dfs(dev, path, allPaths)
			}

		}
	}

	path = path[:len(path)-1]
}

func printp(devs []*device) {
	for _, dev := range devs {
		fmt.Printf("%v,", dev.name)
	}
	fmt.Println()
}

func print(devs []device) {
	for _, dev := range devs {
		fmt.Printf("name: \t%v\n", dev.name)
		fmt.Printf("conns: \n")
		for _, conn := range dev.conns {
			fmt.Printf("\t%s\n", conn.name)
		}
		fmt.Println("------------------------------")
	}
}
