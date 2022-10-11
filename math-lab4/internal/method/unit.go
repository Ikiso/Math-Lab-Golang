package method

import (
	"fmt"
	"github.com/oelmekki/matrix"
	Matrix "math-lab4/pkg/matrix"
)

func Unit(A matrix.Matrix, B []float64, size int) {
	E := matrix.ZeroMatrixFrom(A)
	Result := make([]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			E.SetAt(i, j, 1)
		}
	}

	inverse, _ := Matrix.InverseMatrix(A, size)
	Result, _ = inverse.VectorMultiply(B)

	println("\nРешение СЛАУ методом Единичной матрицы:")
	for i := 0; i < size; i++ {
		fmt.Printf("%v ", Result[i])
	}
}
