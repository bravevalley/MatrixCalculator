package main

import (
	"fmt"
	"math/rand"
)

const matrixSize = 4

var (
	matrixA = [matrixSize][matrixSize]int{}

	matrixB = [matrixSize][matrixSize]int{}

	result [matrixSize][matrixSize]int
)

func preformatRow(row int) {
	for col := 0; col < matrixSize; col++ {
		for i := 0; i < matrixSize; i++ {
			// Formular:
			// [r1c1*r1c1+r1c2*r2c1         r1c1*r1c2+r1c2*r2c2]
			// [r2c1*r1c+r2c2*r2c1          r2c2*r1c2+r2c2*r2c2]

			result[row][col] += matrixA[row][i] * matrixB[i][col]

		}
	}

}

func randMatrixGenerator(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {

			// Assign an element to each index of the matrix
			matrix[row][col] = rand.Intn(10) - 5

		}

	}

}

func main() {

	// Generate random numbers to fill the matrices
	randMatrixGenerator(&matrixA)
	randMatrixGenerator(&matrixB)

	// Pass the row into the row function.
	for row := 0; row < matrixSize; row++ {
		preformatRow(row)

	}

	fmt.Printf("%-v\n", result)
	fmt.Printf("%-v\n", matrixA)
	fmt.Printf("%-v\n", matrixB)

}
