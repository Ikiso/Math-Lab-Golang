package test

import (
	"errors"
	"github.com/oelmekki/matrix"
)

type Test struct {
	FirstMat  matrix.Matrix
	SecondMat matrix.Matrix
	Vector    matrix.Matrix
	Size      int
}

func NewTest(first, second, vector matrix.Matrix, size int) *Test {
	return &Test{
		FirstMat:  first,
		SecondMat: second,
		Vector:    vector,
		Size:      size,
	}
}

// testSolveLU Функция теста решения для lU метода
// используется для проверки совпадения ответов, если
// метод с использование matrix вовзращает не те значения
func (t *Test) testSolveLU(Lower, Upper, vec matrix.Matrix) []float64 {
	n := vec.Cols()
	testy := make([]float64, n)
	testx := make([]float64, n)
	for i := 0; i < n; i++ {
		testy[i] = vec.At(0, i)
		for j := 0; j < i; j++ {
			testy[i] -= Lower.At(i, j) * testy[j]
		}
	}

	for i := n - 1; i >= 0; i-- {
		testx[i] = testy[i]
		for j := i + 1; j < n; j++ {
			testx[i] -= Upper.At(i, j) * testx[j]
		}
		testx[i] /= Upper.At(i, i)
	}

	return testx
}

// TODO: не дописана
func (t *Test) testInverse(A matrix.Matrix, size int) (matrix.Matrix, error) {

	// Копия исходной матрицы (будет использоваться как результирующая)
	Result := matrix.ZeroMatrixFrom(A)
	// Проверяем валидность нашей копии
	if !Result.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}

	// заполняем нашу копию
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Result.SetAt(i, j, A.At(i, j))
		}
	}

	// Единичная матрица
	E := matrix.ZeroMatrixFrom(A)
	// Проверяем валидна ли наша результирующая матрица
	if !E.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}
	// Заполняем единичную матрицу
	for i := 0; i < size; i++ {
		E.SetAt(i, i, 1.0)
	}

	for k := 0; k < size; k++ {
		// Создаем временную переменную для хранения значения
		temp := Result.At(k, k)

		for j := 0; j < size; j++ {
			Result.SetAt(k, j, Result.At(k, j)/temp)
			E.SetAt(k, j, E.At(k, j)/temp)
		}

		for i := k + 1; i < size; i++ {
			temp := Result.At(i, k)

			for j := 0; j < size; j++ {
				Result.SetAt(i, j, Result.At(i, j)-(Result.At(k, j)*temp))
				E.SetAt(i, j, E.At(i, j)-(E.At(k, j)*temp))
			}
		}
	}

	return Result, nil
}
