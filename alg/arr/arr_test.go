package arr

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{0, 2, 3, 4, 5, 6}
	target := 1

	rs := search(nums, target)
	fmt.Println("hello world: ", rs)
}