package main

import (
	"errors"
	"fmt"
	"github.com/oelmekki/matrix"
	"math"
)

func main() {
	// Исходная матрица
	matA, err := matrix.Build(matrix.Builder{
		matrix.Row{15, -7.4, 0, 0, 0},
		matrix.Row{3.8, 11.304, -4.88, 0, 0},
		matrix.Row{0, -5.71, 32.22, -12.6, 0},
		matrix.Row{0, 0, 11, 29.3, 7.5},
		matrix.Row{0, 0, 0, 5.3, -14},
	})
	if err != nil {
		panic(err)
	}
	// Вектор решений
	matB, err := matrix.Build(matrix.Builder{
		matrix.Row{19.64, 10.2496, 37.42, 156.39, -38.51},
	})
	if err != nil {
		panic(err)
	}

	size := matB.Cols()

	if size != matA.Cols() {
		err = errors.New("ошибка кол-во строк матрицы А и вектора Б не равны")
		panic(err)
	}

	err = Running(matA, matB, size)
	if err != nil {
		panic(err)
	}

}

func Running(A, B matrix.Matrix, size int) error {
	// делаем копии наших исходных данных
	// Дальше будем использовать их как исходные
	copyA := matrix.ZeroMatrixFrom(A)
	copyB := matrix.ZeroMatrixFrom(B)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyA.SetAt(i, j, A.At(i, j))
		}
		copyB.SetAt(0, i, B.At(0, i))
	}

	// Матрица ответов СЛАУ
	Result := matrix.ZeroMatrixFrom(B)

	// Создадим ещё несколько матриц диагоналей нашей исходной матрица А
	upper := matrix.ZeroMatrixFrom(B)
	lower := matrix.ZeroMatrixFrom(B)
	mainDiagonal := matrix.ZeroMatrixFrom(B)
	alpha := matrix.ZeroMatrixFrom(B)
	beta := matrix.ZeroMatrixFrom(B)

	for i := 0; i < size; i++ {
		mainDiagonal.SetAt(0, i, copyA.At(i, i))
	}
	for i := 0; i < size-1; i++ {
		lower.SetAt(0, i+1, copyA.At(i+1, i))
		upper.SetAt(0, i, copyA.At(i, i+1))
	}

	for i := 0; i < size; i++ {
		if math.Abs(mainDiagonal.At(0, i)) < math.Abs(lower.At(0, i))+math.Abs(upper.At(0, i)) {
			panic(errors.New("ошибка условие преминимости не выполняется"))
		}
	}

	alpha.SetAt(0, 0, -upper.At(0, 0)/mainDiagonal.At(0, 0))
	beta.SetAt(0, 0, copyB.At(0, 0)/mainDiagonal.At(0, 0))

	for i := 1; i < size; i++ {
		alpha.SetAt(0, i, -upper.At(0, i)/(lower.At(0, i)*alpha.At(0, i-1)+mainDiagonal.At(0, i)))
		beta.SetAt(0, i, (copyB.At(0, i)-lower.At(0, i)*beta.At(0, i-1))/
			(lower.At(0, i)*alpha.At(0, i-1)+mainDiagonal.At(0, i)))
	}

	// Обратный ход где получаем решение СЛАУ
	Result.SetAt(0, size-1, (copyB.At(0, size-1)-lower.At(0, size-1)*beta.At(0, size-2))/
		(lower.At(0, size-1)*alpha.At(0, size-2)+mainDiagonal.At(0, size-1)))
	for i := size - 2; i >= 0; i-- {
		Result.SetAt(0, i, alpha.At(0, i)*Result.At(0, i+1)+beta.At(0, i))
	}

	fmt.Printf("\nВектор решений СЛАУ: %v\n", Result.String())
	estimate := EstimateError(copyA, copyB, Result)
	fmt.Printf("\nПогрешность равна: %.20f\n", estimate)
	return nil
}

func EstimateError(A, vec, x matrix.Matrix) float64 {
	outragedRightSide := getOutragedRightSide(A, x)
	vector := matrix.ZeroMatrixFrom(vec)

	for i := 0; i < x.Cols(); i++ {
		vector.SetAt(0, i, vec.At(0, i)-outragedRightSide.At(0, i))
	}
	return getEuclideanNorm(vector) / getEuclideanNorm(vec)
}

func getOutragedRightSide(A, x matrix.Matrix) matrix.Matrix {
	outragedRightSide := matrix.ZeroMatrixFrom(x)

	for i := 0; i < x.Cols(); i++ {
		for j := 0; j < x.Cols(); j++ {
			temp := x.At(0, j) * A.At(i, j)
			outragedRightSide.SetAt(0, i, outragedRightSide.At(0, i)+temp)
		}
	}
	return outragedRightSide
}

func getEuclideanNorm(vector matrix.Matrix) float64 {
	sum := 0.0
	for i := 0; i < vector.Cols(); i++ {
		vector.SetAt(0, i, vector.At(0, i)*vector.At(0, i))
		sum += vector.At(0, i)
	}
	return math.Pow(sum, 0.5)
}
