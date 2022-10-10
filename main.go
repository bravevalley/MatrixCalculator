package main

import (
	"fmt"
	"math/rand"
	"time"

	"com.github/Barrier"
)

// Define the matrix size
const matrixSize = 4

var (
	// Declara the matricies
	matrixA = [matrixSize][matrixSize]int{}

	matrixB = [matrixSize][matrixSize]int{}

	result [matrixSize][matrixSize]int

	// Create a barrier structures
	barrierOne = barrier.NewBarrier(4)

	barrierTwo = barrier.NewBarrier(4)
)

// This is where the calculation takes place 
// We receive the row index as an argument
func preformatRow(row int) {

	// Start an infinite loop that keep calculating until we have 
	// exhausted the rows
	for {

		// Create a barrier that waits for the main thread to generate the
		// matrices 
		barrierOne.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				// Formular:
				// [r1c1*r1c1+r1c2*r2c1         r1c1*r1c2+r1c2*r2c2]
				// [r2c1*r1c+r2c2*r2c1          r2c2*r1c2+r2c2*r2c2]

				result[row][col] += matrixA[row][i] * matrixB[i][col]

			}
		}

		// Now we wait for the  other threads to reach this barrier as well
		// And the main thread can pump another batch of rows
		barrierOne.Wait()
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

	// Pass the row into the row function.
	for row := 0; row < matrixSize; row++ {
		go preformatRow(row)

	}

	// Generate random numbers to fill the matrices
	for i := 0; i < 100; i++ {
		
		randMatrixGenerator(&matrixA)
		randMatrixGenerator(&matrixB)

		// This first barrier is to wait till worker threads are ready
		barrierOne.Wait()

		// This is to wait for worker threads to finish computation
		barrierTwo.Wait()
	}

	finish := time.Since(start)
	fmt.Println(finish)
}


/*
	The worker thread starts up by entering a wait in their infinite loop
	The main thread then writes the random number to the matricies
	After writing the numbers, the main thread reaches it wait func as well
	And immediately enter another wait()
	The first wait the mainthread entered would trigger the worker threads to 
		begin computation
	The second wait instance of the main thread is to pause until the worker 
		threads are done computing and reach their second wait as well.
	When the worker threads reach their second wait, this would unblock the main
		thread, the main thread repeats the cycle with the worker thread


	Main Thread --load matrix--|wait& --------|wait@-- load matrix REPEAT
	Worker1 -|wait& ----------| compute ----|wait@---- |wait&   REPEAT
	Worker2 -|wait& ----------| compute ----|wait@---- |wait&   REPEAT
	Worker3 -|wait& ----------| compute ----|wait@---- |wait&   REPEAT
	

*/