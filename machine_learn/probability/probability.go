package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	ob := []float64{300, 200, 100}
	total := 600.0
	ex := []float64{total * 0.5, total * 0.3, total * 0.2}
	fmt.Println(stat.ChiSquare(ob, ex))
}
