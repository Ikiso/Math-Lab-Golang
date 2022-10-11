package methods

import (
	"math"
	"math-lab5/pkg/matrix"
	"math-lab5/pkg/normal"
)

func Stepwise(A [][]float64, size int, eps float64) ([]float64, []float64, float64, int) {
	// зададим произвольный не нулевой вектор и вектор нормы
	vector := make([]float64, size)
	norm := make([]float64, size)
	prev := make([]float64, size)
	current := make([]float64, size)

	lambda := 0.0

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
	for i := 0; i < size; i++ {
		prev[i] = vector[i] / norm[i]
	}

	iteration := 2

	for {

		// Нормируем вектор vector
		for i := 0; i < size; i++ {
			norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
		}

		// Вычислим новые значения вектора по формуле y = Ax
		vector, _ = matrix.ScalarProductMatrixToVector(copyA, norm, size, size)

		// Вычисляем лямбду
		for i := 0; i < size; i++ {
			current[i] = vector[i] / norm[i]
		}

		if maxDifference(current, prev) <= eps {

			lambda = findLambda(current)
			// Нормируем вектор vector
			for i := 0; i < size; i++ {
				norm[i] = vector[i] / normal.GetEuclideanNorm(vector)
			}
			break
		} else {
			for i := 0; i < size; i++ {
				prev[i] = current[i]
			}
			iteration++
		}

	}

	return current, norm, lambda, iteration
}

func findLambda(lambda []float64) (value float64) {
	for i := 0; i < len(lambda); i++ {
		value += lambda[i]
	}

	value /= float64(len(lambda))

	return
}

func maxDifference(first, second []float64) (max float64) {
	max = -9999

	for i := 0; i < len(first); i++ {
		if math.Abs(first[i]-second[i]) > max {
			max = math.Abs(first[i] - second[i])
		}
	}

	return
}
