package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(lib.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		challenge.Function("Fibonacci", student.Fibonacci, solutions.Fibonacci, arg)
	}
}
