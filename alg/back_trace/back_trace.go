package flashback

import "fmt"

func Print(a []int, i int) {
	fmt.Print("{")
	for j := 0; j < i; j++ {
		fmt.Print(a[j], ",")
	}
	fmt.Print("}")
}

func Backtrace(a []int, i int, box, answer []int)
