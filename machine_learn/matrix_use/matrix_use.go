package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
)

func main() {
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(mat64.Formatted(a.T()))
	fmt.Println(mat64.Det(a)) //行列式
}
