package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(5, 5))
	//Output: 10
}

func BenchmarkSum(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Sum(1, 2)
	}
}
