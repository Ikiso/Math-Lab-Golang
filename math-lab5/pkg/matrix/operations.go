package matrix

import (
	"errors"
	"github.com/oelmekki/matrix"
	"math"
)

func ScalarProductVectors(firstVector, secondVector []float64) (result float64) {
	if len(firstVector) != len(secondVector) {
		err := errors.New("len vectors not equals")
		panic(err)
	}

	for i := 0; i < len(firstVector); i++ {
		result += firstVector[i] * secondVector[i]
	}

	return
}

func ScalarProductMatrixToVector(Matrix [][]float64, vector []float64, rows, cols int) (result []float64, err error) {
	if len(vector) != rows {
		err = errors.New("count rows matrix not equals len vector")
		return
	}

	result = make([]float64, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i] += Matrix[i][j] * vector[j]
		}
	}

	return
}

// InverseMatrix Получение обратной матрицы матрице А
// size - это размер матрицы (т.к матрица
// квадратная, то можно подставить кол-во либо Rows/Cols
// Если матрица не квадтраная, то будет возвращенна ошибка
func InverseMatrix(A matrix.Matrix, size int) (matrix.Matrix, error) {
	// Проверяем совпадает ли количество Rows и Cols если не совпадает
	// возвращаем ошибку
	if !A.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}
	// Создаем нашу обратную матрицу в виде 0 матрицы исходной (результат обработки)
	Result := matrix.ZeroMatrixFrom(A)
	// Проверяем валидна ли наша результирующая матрица
	if !Result.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}
	// Изначально результирующая матрица является единичной
	// Заполняем единичную матрицу
	for i := 0; i < size; i++ {
		Result.SetAt(i, i, 1.0)
	}

	// Создаем и заполняем единичную матрицу для будущей проверки
	Unit := matrix.ZeroMatrixFrom(A)

	for i := 0; i < size; i++ {
		Unit.SetAt(i, i, 1.0)
	}
	// Проверяем валидна ли наша единичная матрица
	if !Result.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}

	// Копия исходной матрицы
	copyMat := matrix.ZeroMatrixFrom(A)
	// Проверяем валидность нашей копии
	if !copyMat.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		return nil, err
	}

	// заполняем нашу копию
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyMat.SetAt(i, j, A.At(i, j))
		}
	}

	// Проходим по строкам матрицы (назовём их исходными)
	// сверху вниз. На данном этапе происходит прямой ход
	// и исходная матрица превращается в верхнюю треугольную
	for k := 0; k < size; k++ {
		// Если элемент на главной диагонали в исходной
		// строке - нуль, то ищем строку, где элемент
		// того же столбца не нулевой, и меняем строки
		// местами
		if math.Abs(copyMat.At(k, k)) < 0.00000001 {
			// Ключ, говорящий о том, что был произведён обмен строк
			changed := false

			// Идём по строкам, расположенным ниже исходной
			for i := k + 1; i < size; i++ {
				// Если нашли строку, где в том же столбце
				// имеется ненулевой элемент
				if math.Abs(copyMat.At(i, k)) > 0.00000001 {
					// Меняем найденную и исходную строки местами
					// как в исходной матрице, так и в единичной
					swap(copyMat, copyMat.Cols(), i, k)
					swap(Result, Result.Cols(), i, k)
					changed = true
					break
				}
			}
			if !changed {
				err := errors.New("ошибка не возможно получить обратную матрицу")
				return nil, err
			}
		}
		// Запоминаем делитель - диагональный элемент
		div := copyMat.At(k, k)

		// Все элементы исходной строки делим на диагональный
		// элемент как в исходной матрице, так и в единичной
		for j := 0; j < size; j++ {
			copyMat.SetAt(k, j, copyMat.At(k, j)/div)
			Result.SetAt(k, j, Result.At(k, j)/div)
		}

		// Идём по строкам, которые расположены ниже исходной
		for i := k + 1; i < size; i++ {
			// Запоминаем множитель - элемент очередной строки,
			// расположенный под диагональным элементом исходной
			// строки
			multi := copyMat.At(i, k)

			// Отнимаем от очередной строки исходную, умноженную
			// на сохранённый ранее множитель как в исходной,
			// так и в единичной матрице
			for j := 0; j < size; j++ {
				copyMat.SetAt(i, j, copyMat.At(i, j)-(multi*copyMat.At(k, j)))
				Result.SetAt(i, j, Result.At(i, j)-(multi*Result.At(k, j)))
			}
		}
	}

	// Проходим по вернхней треугольной матрице, полученной
	// на прямом ходе, снизу вверх
	// На данном этапе происходит обратный ход, и из исходной
	// матрицы окончательно формируется единичная, а из единичной -
	// обратная
	for k := size - 1; k >= 0; k-- {
		// Идём по строкам, которые расположены выше исходной
		for i := k - 1; i+1 > 0; i-- {
			// Запоминаем множитель - элемент очередной строки,
			// расположенный над диагональным элементом исходной
			// строки
			multi := copyMat.At(i, k)

			// Отнимаем от очередной строки исходную, умноженную
			// на сохранённый ранее множитель как в исходной,
			// так и в единичной матрице
			for j := 0; j < size; j++ {
				copyMat.SetAt(i, j, copyMat.At(i, j)-(multi*copyMat.At(k, j)))
				Result.SetAt(i, j, Result.At(i, j)-(multi*Result.At(k, j)))
			}
		}
	}

	// Проверка что при Result * Unit будет наша исходная матрица А
	temp := matrix.ZeroMatrixFrom(Result)
	temp, _ = Result.DotProduct(A)
	// Т.к есть погрешность, то необходима округление элементов
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			temp.SetAt(i, j, math.Round(temp.At(i, j)))
		}
	}
	if !Unit.EqualTo(temp) {
		err := errors.New("матрицы не совпадают")
		return nil, err
	}

	return Result, nil
}

// Swap функция которая заменят строки местами
// исходной матрицы А типа matrix.Matrix, принимает параметры исходную матрицу типа matrix.Matrix
// ,размер, индексы заменяемых строк типа int, возвращаем изминённую матрицу А
// типа matrix.Matrix
func swap(A matrix.Matrix, size, firstLineIndex, secondLineIndex int) matrix.Matrix {
	//A[i][j], A[j][i] = A[j][i], A[i][j]

	// Производим замену строк
	for k := 0; k < size; k++ {
		temp := A.At(firstLineIndex, k)
		A.SetAt(firstLineIndex, k, A.At(secondLineIndex, k))
		A.SetAt(secondLineIndex, k, temp)
	}
	return A
}
