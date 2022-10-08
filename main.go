package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 4

var (
	matrixA = [matrixSize][matrixSize]int{}

	matrixB = [matrixSize][matrixSize]int{}

	result [matrixSize][matrixSize]int

	rwLock = sync.RWMutex{}

	condition = sync.NewCond(rwLock.RLocker())

	wg = sync.WaitGroup{}
)

func preformatRow(row int) {

	rwLock.RLock()
	for {
		wg.Done()
		condition.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				// Formular:
				// [r1c1*r1c1+r1c2*r2c1         r1c1*r1c2+r1c2*r2c2]
				// [r2c1*r1c+r2c2*r2c1          r2c2*r1c2+r2c2*r2c2]

				result[row][col] += matrixA[row][i] * matrixB[i][col]

			}
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

	start := time.Now()
	wg.Add(matrixSize)

	// Pass the row into the row function.
	for row := 0; row < matrixSize; row++ {
		go preformatRow(row)

	}

	// Generate random numbers to fill the matrices
	for i := 0; i < 100; i++ {
		wg.Wait()
		rwLock.Lock()
		randMatrixGenerator(&matrixA)
		randMatrixGenerator(&matrixB)
		wg.Add(matrixSize)
		rwLock.Unlock()
		condition.Broadcast()
	}

	finish := time.Since(start)
	fmt.Println(finish)
}
