package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int64{0, 1, 3, 5, 6, 7, 8, 9, 10}

	fmt.Println(BinarySearch(a, -1))
}