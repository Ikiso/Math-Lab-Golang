package method

import (
	"fmt"
	"math"
	"math-lab4/pkg/matrix"
)

func Diagonal(A [][]float64, size int) {
	for i := 0; i < size; i++ {
		for k := i + 1; k < size; k++ {
			if math.Abs(A[i][i]) < math.Abs(A[k][i]) {
				for j := 0; j <= size; j++ {
					temp := A[i][j]
					A[i][j] = A[k][j]
					A[k][j] = temp
				}
			}
		}
	}
}

func Seidel(A [][]float64, B []float64, size int, eps float64) {
	// генерируем результирующую матрицу и создаем копии ветора и исходной матрицы
	resultMatrix := make([]float64, size)
	copyA := make([][]float64, len(A))
	for i := range copyA {
		copyA[i] = make([]float64, len(A[i]))
	}
	copyB := make([]float64, len(B))
	for i := 0; i < size; i++ {
		for j := 0; j <= size; j++ {
			copyA[i][j] = A[i][j]
		}
		copyB[i] = B[i]
	}

	k := 0

	Diagonal(copyA, size)
	prevX := 0.0
	flag := false
	// метод Зейделя
	for {
		k++
		for i := 0; i < size; i++ {
			prevX = resultMatrix[i]
			resultMatrix[i] = copyA[i][size]
			for j := 0; j < size; j++ {
				if i != j {
					resultMatrix[i] = resultMatrix[i] - copyA[i][j]*resultMatrix[j]
				}
			}
			resultMatrix[i] = resultMatrix[i] / copyA[i][i]
			if math.Abs(resultMatrix[i]-prevX) < eps {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	printResult(resultMatrix, k, eps)
}

func printResult(result []float64, iteration int, eps float64) {
	fakeMatrix := matrix.Build(result)
	fmt.Printf("\nКол-во итераций: %d,Точность: %e, Решение СЛАУ: %v",
		iteration, eps, fakeMatrix)
}
