package methods

import (
	"math"
	"math-lab5/pkg/matrix"
	"math-lab5/pkg/normal"
)

func FindMaxAndMin(A [][]float64, size int, eps float64) (max, min float64) {
	copyA := matrix.BuildMatrixFloat64(A, size, size)
	fakeMatrix := matrix.ConvertMatrixToMatrix(copyA)

	inverse, _ := matrix.InverseMatrix(fakeMatrix, size)

	mat := matrix.ConvertMatrixToFloat64(inverse)

	max = findMin(mat, size, eps)
	min = findMax(copyA, size, eps)

	return
}

func findMax(A [][]float64, size int, eps float64) float64 {
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
	prev := matrix.ScalarProductVectors(vector, norm) / matrix.ScalarProductVectors(norm, norm)

	iteration := 1

	for {

		// Нормируем вектор vector
		for i := 0; i < size; i++ {
			norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
		}

		// Вычислим новые значения вектора по формуле y = Ax
		vector, _ = matrix.ScalarProductMatrixToVector(copyA, norm, size, size)

		// Вычисляем лямбду
		current = matrix.ScalarProductVectors(vector, norm) / matrix.ScalarProductVectors(norm, norm)

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
	return current
}

func findMin(A [][]float64, size int, eps float64) float64 {
	// зададим произвольный не нулевой вектор и вектор нормы
	vector := make([]float64, size)
	norm := make([]float64, size)

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
	μ := matrix.ScalarProductVectors(vector, norm) / matrix.ScalarProductVectors(vector, vector)
	prev := 1 / μ
	current := 0.0

	iteration := 2

	for {

		// Нормируем вектор vector
		for i := 0; i < size; i++ {
			norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
		}

		// Вычислим новые значения вектора по формуле y = Ax
		vector, _ = matrix.ScalarProductMatrixToVector(copyA, norm, size, size)

		// Вычисляем лямбду
		μ = matrix.ScalarProductVectors(vector, norm) / matrix.ScalarProductVectors(vector, vector)
		current = 1 / μ

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
	return current
}
