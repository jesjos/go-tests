package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func drawLine(x int, s string) {
	beg := s[0]
	med := s[1]
	end := s[2]
	if x >= 1 {
		fmt.Printf("%c", beg)
	}
	if x > 2 {
		for i := 0; i < (x - 2); i++ {
			fmt.Printf("%c", med)
		}
	}
	if x > 1 {
		fmt.Printf("%c", end)
	}
	fmt.Println()
}

func printTheLines(x, y int, strBeg, strMed, strEnd string) {
	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func quadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	printTheLines(x, y, "o-o", "| |", "o-o")
}

func quadB(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "/*\\"
	strMed := "* *"
	strEnd := "\\*/"

	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func quadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "ABA"
	strMed := "B B"
	strEnd := "CBC"

	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func quadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "ABC"
	strMed := "B B"
	strEnd := "ABC"

	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func quadE(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	strBeg := "ABC"
	strMed := "B B"
	strEnd := "CBA"

	if y >= 1 {
		drawLine(x, strBeg)
	}
	if y > 2 {
		for i := 0; i < y-2; i++ {
			drawLine(x, strMed)
		}
	}
	if y > 1 {
		drawLine(x, strEnd)
	}
}

func main() {
	// testing examples of subjects
	table := []int{
		5, 3,
		5, 1,
		1, 1,
		1, 5,
	}

	// testing special cases and one valid random case.
	table = append(table,
		0, 0,
		-1, 6,
		6, -1,
		random.IntBetween(1, 20), random.IntBetween(1, 20),
	)

	// Tests all possibilities including 0 0, -x y, x -y
	for i := 0; i < len(table); i += 2 {
		if i != len(table)-1 {
			challenge.Function("quadA", quadA, student.QuadA, table[i], table[i+1])
			challenge.Function("quadB", quadB, student.QuadB, table[i], table[i+1])
			challenge.Function("quadC", quadC, student.QuadC, table[i], table[i+1])
			challenge.Function("quadD", quadD, student.QuadD, table[i], table[i+1])
			challenge.Function("quadE", quadE, student.QuadE, table[i], table[i+1])
		}
	}
}
