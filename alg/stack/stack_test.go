package stack

import (
	"fmt"
	"testing"
)

func longestValidParentheses(s string) int {
	var tmpLen int
	var maxLen int

	leftBrackets := 0
	for i, v := range s {
		if string(v) == "(" {
			leftBrackets++
			tmpLen++
		} else {
			if leftBrackets > 0 {
				leftBrackets--
				tmpLen++
			} else {
				tmpLen = 0
			}
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}

	return maxLen
}

func TestLongesValidParentheses(t *testing.T) {
	s := "(((((()()))))))))"
	maxLen := longestValidParentheses(s)
	fmt.Println("maxLen : ", maxLen)
}
