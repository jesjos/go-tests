package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, solutions.PrintStr, arg)
	}
}
