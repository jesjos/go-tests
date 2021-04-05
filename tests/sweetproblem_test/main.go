package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		red   int
		green int
		blue  int
	}

	table := []node{
		{50, 43, 20},
		{10, 0, 0},
		{0, 10, 0},
		{0, 0, 10},
		{10, 1, 0},
		{0, 10, 1},
		{1, 0, 10},
		{10, 2, 0},
		{2, 10, 0},
		{0, 2, 10},
		{13, 13, 0},
		{10, 9, 0},
		{5, 9, 2},
	}

	for i := 0; i < 15; i++ {
		table = append(table, node{
			red:   lib.RandIntBetween(0, 30),
			green: lib.RandIntBetween(0, 30),
			blue:  lib.RandIntBetween(0, 30),
		})
	}
	for _, arg := range table {
		challenge.Function("SweetProblem", student.Sweetproblem, solutions.Sweetproblem, arg.red, arg.green, arg.blue)
	}
}
