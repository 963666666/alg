package test

import (
	"fmt"
	"testing"
)

func TestFort(t *testing.T) {
	if 1 == 1 {
		fmt.Println("hello world 1==1")
	} else if 2 == 2 {
		fmt.Println("hello world 2==2")
	} else {
		fmt.Println("hello world 3==3")
	}
}
