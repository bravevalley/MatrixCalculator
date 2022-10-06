package main

import "fmt"

var (
	matrixSize = 3
	matrixA    = [][]int{{9, -4, -6},
		{-5, 3, 2},
		{1, -3, 2}}

	matrixB = [][]int{{9, -3, 7},
		{4, 1, -2},
		{6, 5, -4}}
)

func main() {
	fmt.Printf("%T", matrixA)
	fmt.Printf("%T", matrixB)

}
