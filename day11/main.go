package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

type SafeMap struct {
	data map[string][][]*device
	mu   sync.RWMutex // The Read/Write Mutex
}

// Write operation: Requires an exclusive Lock()
func (sm *SafeMap) Set(key string, value []*device) {
	sm.mu.Lock()         // Acquire an exclusive lock for writing
	defer sm.mu.Unlock() // Ensure the lock is released when the function exits

	sm.data[key] = append(sm.data[key], value)
}

// Read operation: Requires a shared RLock()
func (sm *SafeMap) Get(key string) ([][]*device, bool) {
	sm.mu.RLock()         // Acquire a shared read lock
	defer sm.mu.RUnlock() // Ensure the read lock is released

	val, ok := sm.data[key]
	return val, ok
}

func (sm *SafeMap) Len() int {
	sm.mu.RLock()         // Acquire a shared read lock
	defer sm.mu.RUnlock() // Ensure the read lock is released

	return len(sm.data)
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string][][]*device)}
}

var cache *SafeMap

func main() {
	cache = NewSafeMap()
	input := read()
	devs := parse(input)
	// print(devs)
	allPaths := make([][]*device, 0)
	you := slices.IndexFunc(devs, func(d *device) bool {
		return d.name == "svr"
		// return d.name == "you"
	})

	dfs(devs[you], make([]*device, 0), &allPaths)
	// alg(devs, &allPaths)
	fmt.Printf("you: %v\n", you)

	// printAp(allPaths)

	cnt := 0
	for _, path := range allPaths {
		b1 := slices.ContainsFunc(path, func(d *device) bool {
			return d.name == "fft"
		})

		b2 := slices.ContainsFunc(path, func(d *device) bool {
			return d.name == "dac"
		})

		printp(path)
		if b1 && b2 {
			cnt++
		}
	}

	fmt.Printf("Result: %v\n", len(allPaths))
	fmt.Printf("cnt: %v\n", cnt)
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
	for i, paths := range cache.data {
		fmt.Printf("key %s: ", i)
		for _, path := range paths {
			for _, dev := range path {
				fmt.Printf("%v,", dev.name)
			}
			fmt.Printf("\n\t ")
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

func parse(input []string) []*device {
	ret := make([]*device, 0, len(input))
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
		ret = append(ret, dev)
	}

	slices.SortFunc(ret, func(a, b *device) int {
		return strings.Compare(a.name, b.name)
	})
	return ret
}

func dfs(src *device, path []*device, allPaths *[][]*device) {
	// fmt.Printf("On device: \t %v options:", src.name)
	// for _, conn := range src.conns {
	// 	fmt.Printf("%s,", conn.name)
	// }
	// fmt.Println()

	if slices.ContainsFunc(path, func(d *device) bool {
		return src.name == d.name
	}) {
		fmt.Println("------------------------------Found cicle------------------------------")
		return
	}

	path = append(path, src)

	if src.IsOut() {
		*allPaths = append(*allPaths, path)
		// fmt.Println("-----------------------------Found Out-----------------------------")
	} else {
		for _, dev := range src.conns {
			dfs(dev, path, allPaths)
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

func print(devs []*device) {
	for _, dev := range devs {
		fmt.Printf("name: \t%v\n", dev.name)
		fmt.Printf("conns: \n")
		for _, conn := range dev.conns {
			fmt.Printf("\t%s\n", conn.name)
		}
		fmt.Println("------------------------------")
	}
}

func alg(devices []*device, allPaths *[][]*device) {
	lastLen := 0

	for _, dev := range devices {
		ok := slices.ContainsFunc(dev.conns, func(dev *device) bool {
			return dev.IsOut()
		})

		if ok {
			fmt.Println("Found out")
			cache.Set(dev.name, []*device{dev, {name: "out"}})
		}
	}
	i := 1

	for {
		if lastLen == cache.Len() {
			break
		}

		lastLen = cache.Len()

		fmt.Printf("Passtrought count %d\n", i)
		i++

		for _, dev := range devices {
			devList, _ := cache.Get(dev.name)
			for _, child := range dev.conns {
				if slices.ContainsFunc(devList, func(d []*device) bool {
					return child.name == d[1].name || child.IsOut()
				}) {
					// fmt.Printf("skipping name: %v\n", child.name)
					continue
				}

				paths, ok := cache.Get(child.name)
				if ok {
					for _, path := range paths {
						cache.Set(dev.name, append([]*device{dev}, path...))
					}
				}
			}

		}
		fmt.Printf("cache: %v\n", len(cache.data))
	}

	dev := devices[len(devices)-1]
	for _, child := range dev.conns {
		paths, ok := cache.Get(child.name)
		if ok {
			for _, path := range paths {
				cache.Set(dev.name, append([]*device{dev}, path...))
			}
		}
	}
	data, _ := cache.Get("you")
	*allPaths = append(*allPaths, data...)
}
