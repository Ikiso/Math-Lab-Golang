package methods

import (
	"math"
	"math-lab5/pkg/matrix"
	"math-lab5/pkg/normal"
)

func ScalarProducts(A [][]float64, size int, eps float64) ([]float64, float64, int) {
	// зададим произвольный не нулевой вектор и вектор нормы
	vector := make([]float64, size)
	norm := make([]float64, size)

	current := 0.0

	for i := 0; i < size; i++ {
		vector[i] = 1
	}

	// Делаем копию нашей исходной матрицы
	copyA := matrix.BuildMatrixFloat64(A, size, size)

	// Нормируем вектор vector
	for i := 0; i < size; i++ {
		norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
	}
	// Вычислим новые значения вектора по формуле y = Ax
	vector, _ = matrix.ScalarProductMatrixToVector(copyA, norm, size, size)

	// Вычисляем лямбду
	prev := matrix.ScalarProductVectors(vector, vector) / matrix.ScalarProductVectors(vector, norm)

	iteration := 2

	for {

		// Нормируем вектор vector
		for i := 0; i < size; i++ {
			norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
		}

		// Вычислим новые значения вектора по формуле y = Ax
		vector, _ = matrix.ScalarProductMatrixToVector(copyA, norm, size, size)

		// Вычисляем лямбду
		current = matrix.ScalarProductVectors(vector, vector) / matrix.ScalarProductVectors(vector, norm)

		if math.Abs(current-prev) <= eps {
			// Нормируем вектор vector
			for i := 0; i < size; i++ {
				norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
			}
			break
		} else {
			prev = current
			iteration++
		}
	}

	return norm, current, iteration
}
