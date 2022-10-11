package main

import (
	"github.com/oelmekki/matrix"
	"math-lab4/internal/method"
)

func init() {
	accuracys := [6]float64{0.001, 0.0001, 0.00001, 0.000001, 0.0000001, 0.00000001}
	size := 4

	mat := [][]float64{
		{8, 11, -1, -0.07, 0.654},
		{4.056, 0.3, -5.3, 0.11, -24.016},
		{0.2, 5.77, 2.8, -15, 54.567},
		{12, -3.8, 1, 3, -64.78},
	}

	A, _ := matrix.Build(matrix.Builder{
		matrix.Row{8, 11, -1, -0.07},
		matrix.Row{4.056, 0.3, -5.3, 0.11},
		matrix.Row{0.2, 5.77, 2.8, -15},
		matrix.Row{12, -3.8, 1, 3},
	})

	vec := []float64{0.654, -24.016, 54.567, -64.78}

	for i := 0; i < 6; i++ {
		method.Seidel(mat, vec, size, accuracys[i])
	}
	method.Unit(A, vec, size)
}

func main() {

}
