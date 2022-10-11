package matrix

import (
	"errors"
	"fmt"
	"github.com/oelmekki/matrix"
	"math"
	"math/rand"
	"time"
)

const (
	min = -3400.00
	max = 3400.00
)

func randFloat64() float64 {
	return min + rand.Float64()*(max-min)
}

func Generate(size int) (resultMatrix, resultVector matrix.Matrix, err error) {

	resultMatrix = matrix.GenerateMatrix(size, size)
	rand.Seed(time.Now().UnixNano())

	// генерация матрицы и заполнения её рандомными числами
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			resultMatrix.SetAt(i, j, randFloat64())
			if resultMatrix.At(i, j) == math.NaN() {
				err = GenerateError(fmt.Sprintf("ошибка генерации матрицы"))
				return
			}
		}
	}

	// генерация вектора решений матрицы
	resultVector = matrix.GenerateMatrix(size, 1)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		resultVector.SetAt(i, 0, randFloat64())
		if resultVector.At(i, 0) == math.NaN() {
			err = GenerateError(fmt.Sprintf("ошибка генерации вектора"))
			return
		}
	}

	if resultVector.Rows() != resultMatrix.Rows() {
		err = GenerateError(fmt.Sprintf("ошибка кол-во столбцов ветора != кол-ву строк матрицы"))
		return
	}

	return
}

func GestationToTheNormal(A matrix.Matrix, B []float64, size int) (resultMatrix matrix.Matrix, resultVector []float64) {

	fakeMatrix := matrix.ZeroMatrixFrom(A)
	resultVector = make([]float64, size)
	copyA := matrix.ZeroMatrixFrom(A)
	copyB := make([]float64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyA.SetAt(i, j, A.At(i, j))
			fakeMatrix.SetAt(i, j, A.At(i, j))
		}
		copyB[i] = B[i]
		resultVector[i] = B[i]
	}

	inverse, err := InverseMatrix(copyA, size)
	if err != nil {
		panic(err)
	}

	resultMatrix = generationMatrixDiagonalPredominance(size)

	fakeMatrix, _ = resultMatrix.DotProduct(inverse)
	resultVector, _ = fakeMatrix.VectorMultiply(copyB)
	return
}

func generationMatrixDiagonalPredominance(size int) (resultMatrix matrix.Matrix) {
	resultMatrix = matrix.GenerateMatrix(size, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			resultMatrix.SetAt(i, j, 1)
		}
		resultMatrix.SetAt(i, i, 4)
	}
	return
}

// ConvertVectorToFloat64 преобразует передаваемый вектор типа matrix.Matrix
// в вектор типа []float64
func ConvertVectorToFloat64(vector matrix.Matrix) (resultVector []float64, err error) {
	if vector.Cols() != 0 {
		resultVector = buildVectorFloat64(vector, vector.Cols(), "cols")
	} else if vector.Rows() != 0 {
		resultVector = buildVectorFloat64(vector, vector.Rows(), "rows")
	} else {
		err = errors.New("cols and rows vector equal 0")
		return
	}

	return
}

func buildVectorFloat64(vector matrix.Matrix, size int, str string) (resultVector []float64) {
	resultVector = make([]float64, size)

	switch str {
	case "cols":
		for i := 0; i < size; i++ {
			resultVector[i] = vector.At(0, i)
		}
	case "rows":
		for i := 0; i < size; i++ {
			resultVector[i] = vector.At(i, 0)
		}
	}

	return
}

func BuildMatrixFloat64(Matrix [][]float64, rows, cols int) (resultMatrix [][]float64) {
	resultMatrix = make([][]float64, rows)

	for i := range resultMatrix {
		resultMatrix[i] = make([]float64, cols)
	}

	resultMatrix = fillMatrix(resultMatrix, Matrix, rows, cols)

	return
}

func fillMatrix(resultMatrix, Matrix [][]float64, rows, cols int) [][]float64 {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			resultMatrix[i][j] = Matrix[i][j]
		}
	}

	return resultMatrix
}

// ConvertMatrixToFloat64 преобразует передаваемую матрицу типа matrix.Matrix
// в матрицу типа [][]float64
func ConvertMatrixToFloat64(Matrix matrix.Matrix) (resultMatrix [][]float64) {
	resultMatrix = make([][]float64, Matrix.Rows())

	for i := range resultMatrix {
		resultMatrix[i] = make([]float64, Matrix.Cols())
	}

	for i := 0; i < Matrix.Rows(); i++ {
		for j := 0; j < Matrix.Cols(); j++ {
			resultMatrix[i][j] = Matrix.At(i, j)
		}
	}

	return
}

// ConvertVectorToMatrix преобразует передаваемый вектор типа []float64,
// в matrix.Matrix
func ConvertVectorToMatrix(vector []float64) (resultVector matrix.Matrix) {
	resultVector = matrix.GenerateMatrix(len(vector), 1)
	for i := 0; i < len(vector); i++ {
		resultVector.SetAt(i, 0, vector[i])
	}
	return
}

// ConvertMatrixToMatrix преобразует передаваемую матрицу типа [][]float64
// в matrix.Matrix
func ConvertMatrixToMatrix(Matrix [][]float64) (resultMatrix matrix.Matrix) {
	resultMatrix = matrix.GenerateMatrix(len(Matrix), len(Matrix))
	for i := 0; i < len(Matrix); i++ {
		for j := 0; j < len(Matrix); j++ {
			resultMatrix.SetAt(i, j, Matrix[i][j])
		}
	}
	return
}
