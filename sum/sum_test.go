package sum

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

func TestSumAll(t *testing.T) {
	t.Run("[1, 2, 3, 4, 5]", testSumAllFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("[1, 2, 3, 4, -5]", testSumAllFunc([]int{1, 2, 3, 4, -5}, 5))
}

func testSumAllFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SumAll(numbers)
		if actual != expected {
			t.Errorf("Expected the sum of %v to be %d but instead got %d", numbers, expected, actual)
		}
	}
}
