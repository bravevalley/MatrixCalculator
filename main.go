package main

import "fmt"

const matrixSize = 3

var (
	matrixA = [matrixSize][matrixSize]int{{9, -4, -6},
		{-5, 3, 2},
		{1, -3, 2}}

	matrixB = [matrixSize][matrixSize]int{{9, -3, 7},
		{4, 1, -2},
		{6, 5, -4}}
	result [matrixSize][matrixSize]int
)

func main() {

	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				// Formular:
				// [r1c1*r1c1+r1c2*r2c1         r1c1*r1c2+r1c2*r2c2]
				// [r2c1*r1c+r2c2*r2c1          r2c2*r1c2+r2c2*r2c2]
				

				result[row][col] += matrixA[row][i] * matrixB[i][col]

			}
		}
	}

	fmt.Printf("%-v\n", result)

}
