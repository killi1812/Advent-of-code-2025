package main

import (
	"os"
	"strings"
)

func main() {
	input := read()
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
