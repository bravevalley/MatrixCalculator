package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_randMatrixGenerator(t *testing.T) {
	type matrices struct {
		matrix [matrixSize][matrixSize]int
	}

	tests := []matrices{
		{[matrixSize][matrixSize]int{}},
		{[matrixSize][matrixSize]int{}},
		{[matrixSize][matrixSize]int{}},
		{[matrixSize][matrixSize]int{}},
		{[matrixSize][matrixSize]int{}},
		{[matrixSize][matrixSize]int{}},
	}

	for _, v := range tests {
		randMatrixGenerator(&v.matrix)
	}

	for _, value := range tests {
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {
				got := fmt.Sprintf("%T", value.matrix[i][j])
				want := "int"

				if !strings.Contains(got, want) {
					t.Errorf("Want %v, Got %s", want, got)
				}
			}
		}
	}
}

func Test_preformatRow(t *testing.T) {
	var result [matrixSize][matrixSize]int

	matrixA := [matrixSize][matrixSize]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	matrixB := [matrixSize][matrixSize]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	for i := 0; i < matrixSize; i++ {
		func(row int) {

			// matrices

			for col := 0; col < matrixSize; col++ {
				for j := 0; j < matrixSize; j++ {
					// Formular:
					// [r1c1*r1c1+r1c2*r2c1         r1c1*r1c2+r1c2*r2c2]
					// [r2c1*r1c+r2c2*r2c1          r2c2*r1c2+r2c2*r2c2]

					result[row][col] += matrixA[row][j] * matrixB[j][col]

				}
			}

		}(i)
	}

	for index, col := range result {
		for q := 0; q < len(result); q++ {
			if index == 0 && q == 0 && col[q] != 6 {
				t.Error("Error: test failed")
			} else if index == 1 && q == 1 && col[q] != 12 {
				t.Error("Error: test failed")
			} else if index == 2 && q == 2 && col[q] != 18 {
				t.Error("Error: test failed")
			}
		}
	}

}
