package main

import (
	"reflect"
	"sort"

	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func sortIntegerTable(a []int) {
	sort.Ints(a)
}

func main() {
	i := 0
	for i < lib.SliceLen {
		table1 := lib.MultRandIntBetween(-100, 100)

		tableCopyBefore := make([]int, len(table1))
		copy(tableCopyBefore, table1)

		table2 := make([]int, len(table1))
		copy(table2, table1)

		student.SortIntegerTable(table1)
		sortIntegerTable(table2)
		if !reflect.DeepEqual(table1, table2) {
			challenge.Fatalf("SortIntegerTable(%v), table1 == %v instead of %v ", tableCopyBefore, table1, table2)
		}
		i++
	}
}
