package main

import "testing"

func Test_randMatrixGenerator(t *testing.T) {
	type args struct {
		matrix *[matrixSize][matrixSize]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randMatrixGenerator(tt.args.matrix)
		})
	}
}
