package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot/plotter"
	"os"
)

func readData() error {
	path := ""
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	for _, colName := range df.Names() {
		fmt.Println(colName)
		val := make(plotter.Values, df.Nrow())
		for i, floatVal := range df.Col(colName).Float() {
			fmt.Printf("i : %d, float val : %v\n", i, floatVal)
		}
	}
}
