package main

import (
	"errors"
	"fmt"
	"github.com/oelmekki/matrix"
	"math"
	"os"
)

/* Работа использует сторонюю библиотеку github.com/oelmekki/matrix
для работы с матрицами.*/

func main() {

	// Наши исходные данные
	mat, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{8, 11, -1, -0.07},
			matrix.Row{4.056, 0.3, -5.3, 0.11},
			matrix.Row{0.2, 5.77, 2.8, -15},
			matrix.Row{12, -3.8, 1, 3},
		})

	if !mat.Valid() {
		err := errors.New("ошибка кол-во строк и кол-во столбцов не совпадает")
		if err != nil {
			println(err)
			return
		}
	}

	vec, _ := matrix.Build(matrix.Builder{
		matrix.Row{0.654, -24.016, 54.567, -64.78},
	})

	inv, err := inverseMatrix(mat, 4)
	if err != nil {
		println(err)
		os.Exit(2)
	}
	println("Обратная матрица")
	println(inv.String())

	// Вызов методов решения СЛАУ
	println("Метод LU-разложений")
	LUdesm(mat, vec)
	println("\nМетод Гаусса")
	g, err := Gauss(mat, vec)
	if err != nil {
		println(err)
		panic(err)
	}
	println("\nВектор решений СЛАУ методом Гаусса с выбором опорного элемента")
	println(g.String())

}

// inverseMatrix Получение обратной матрицы матрице А
// size - это размер матрицы (т.к матрица
// квадратная, то можно подставить кол-во либо Rows/Cols
// Если матрица не квадтраная, то будет возвращенна ошибка
func inverseMatrix(A matrix.Matrix, size int) (matrix.Matrix, error) {
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
					Swap(copyMat, copyMat.Cols(), i, k)
					Swap(Result, Result.Cols(), i, k)
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
func Swap(A matrix.Matrix, size, firstLineIndex, secondLineIndex int) matrix.Matrix {
	//A[i][j], A[j][i] = A[j][i], A[i][j]

	// Производим замену строк
	for k := 0; k < size; k++ {
		temp := A.At(firstLineIndex, k)
		A.SetAt(firstLineIndex, k, A.At(secondLineIndex, k))
		A.SetAt(secondLineIndex, k, temp)
	}
	return A
}

// Gauss функция возвращающее решение СЛАУ методом Гаусса
// с выбором опорного элемента
// принимает на вход матрицу А matrix.Matrix, вектор vec
// matrix.Matrix, который содержит в себе значения после =.
// Вовзращает веторе решений matrix.Matrix и ошибку.
func Gauss(A, vec matrix.Matrix) (matrix.Matrix, error) {
	// Получаем размер матрицы, т.к матрица квадратная то достаточно получить
	// либо количество столбцов либо количество строк
	size := A.Cols()
	// Создаем наш вектор решений
	Result := matrix.ZeroMatrixFrom(vec)
	// Делаем копии наших исходных данных в дальнейшем будет использоватсья как оригинал
	copyA := matrix.ZeroMatrixFrom(A)
	copyVec := matrix.ZeroMatrixFrom(vec)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyA.SetAt(i, j, A.At(i, j))
			copyVec.SetAt(0, j, vec.At(0, j))
		}
	}

	// Решаем нашу матрицу
	for i := 0; i < size-1; i++ {
		SortRow(&copyA, &copyVec, size, i)

		for j := i + 1; j < size; j++ {
			if copyA.At(i, i) != 0 {
				multi := copyA.At(j, i) / copyA.At(i, i)

				for k := i; k < size; k++ {
					copyA.SetAt(j, k, copyA.At(j, k)-(copyA.At(i, k)*multi))
				}

				copyVec.SetAt(0, j, copyVec.At(0, j)-(copyVec.At(0, i)*multi))
			}
		}
	}
	//ищем решение
	for i := size - 1; i >= 0; i-- {
		Result.SetAt(0, i, copyVec.At(0, i))
		for j := size - 1; j > i; j-- {
			Result.SetAt(0, i, Result.At(0, i)-(copyA.At(i, j)*Result.At(0, j)))
		}

		if copyA.At(i, i) == 0 {
			if Result.At(0, i) == 0 {
				err := errors.New("имеет множество решений")
				return nil, err
			} else {
				err := errors.New("нет решений")
				return nil, err
			}
		}

		Result.SetAt(0, i, Result.At(0, i)/copyA.At(i, i))
	}

	// Детерминант
	d := determinant(copyA, size)
	fmt.Printf("\nДетерминант метод Гаусса: %f\n", d)

	// Оценка ошибки
	mistake := EstimateError(A, vec, Result)
	fmt.Printf("Погрешность для метода Гаусса %.20f в процентах", mistake*100)

	return Result, nil
}

// SortRow фунция для сортировки строк для метода Гаусса
func SortRow(A, vec *matrix.Matrix, size, sortIndex int) {
	maxElement := A.At(sortIndex, sortIndex)
	maxElementIndex := sortIndex

	for i := sortIndex + 1; i < size; i++ {
		if A.At(i, sortIndex) > maxElement {
			maxElement = A.At(i, sortIndex)
			maxElementIndex = i
		}
	}

	//теперь найден максимальный элемент ставим его на верхнее место
	if maxElementIndex > sortIndex { // если это не первый элемент
		//TODO: заменить на swap()
		temp := vec.At(0, maxElementIndex)
		vec.SetAt(0, maxElementIndex, vec.At(0, sortIndex))
		vec.SetAt(0, sortIndex, temp)
		for i := 0; i < size; i++ {
			temp = A.At(maxElementIndex, i)
			A.SetAt(maxElementIndex, i, A.At(sortIndex, i))
			A.SetAt(sortIndex, i, temp)
		}
	}
}

// LUdesm Метод ЛУ-разложения для решения СЛАУ
func LUdesm(mat, vec matrix.Matrix) {
	// Объявляем новые две матрицы upper and lower, как 0 матрицы исходной
	Upper := matrix.ZeroMatrixFrom(mat)
	Lower := matrix.ZeroMatrixFrom(mat)
	// Выситываем значение длины исходной матрицы
	for i := 0; i < mat.Rows(); i++ {
		// Получаем верхнюю матрицу
		for k := i; k < mat.Cols(); k++ {
			// Сумма L(i,j) * U(j,k)
			sum := 0.0
			for j := 0; j < i; j++ {
				sum += Lower.At(i, j) * Upper.At(j, k)
			}
			// Оценка верхней
			Upper.SetAt(i, k, mat.At(i, k)-sum)
		}
		// Получаем нижнию матрицу
		for k := i; k < mat.Cols(); k++ {
			if i == k {
				Lower.SetAt(i, i, 1) // Диагонал = 1
			} else {
				// Сумма L(k,j) * U(j,i)
				sum := 0.0
				for j := 0; j < i; j++ {
					sum += Lower.At(k, j) * Upper.At(j, i)
				}
				// Оценка нижний матрицы
				Lower.SetAt(k, i, (mat.At(k, i)-sum)/Upper.At(i, i))
			}
		}
	}

	// Делаем копии наших исходных данных
	copyA := matrix.ZeroMatrixFrom(mat)
	copyVec := matrix.ZeroMatrixFrom(vec)

	for i := 0; i < copyA.Rows(); i++ {
		for j := 0; j < copyA.Cols(); j++ {
			copyA.SetAt(i, j, mat.At(i, j))
			copyVec.SetAt(0, j, vec.At(0, j))
		}
	}

	// Наёдем решение системы
	x := SolveLU(Lower, Upper, vec)
	fmt.Println("Вектор решений системы СЛАУ методом LU")
	fmt.Println(x.String())

	// Детерминант для метода LU
	n := Upper.Cols()
	det := determinant(Upper, n)
	fmt.Printf("Детерминант СЛАУ LU равен: %f\n", det)

	// Погрешность
	mistake := EstimateError(copyA, copyVec, x)
	fmt.Printf("Погрешность для метода LU-разложений %.20f в процентах", mistake*100)
}

// determinant Метод определения детерминанта функции
// Входные параметры, матрица А matrix.Matrix, и
// размер матрицы n float64.
// Выходные параметры, значение детерминанта float64
func determinant(A matrix.Matrix, n int) float64 {
	det := 0.0
	subMatrix := matrix.ZeroMatrixFrom(A)
	if n == 2 {
		return (A.At(0, 0) * A.At(1, 1)) - (A.At(1, 0) * A.At(0, 1))
	} else {
		for x := 0; x < n; x++ {
			sbi := 0
			for i := 1; i < n; i++ {
				sbj := 0
				for j := 0; j < n; j++ {
					if j == x {
						continue
					}
					subMatrix.SetAt(sbi, sbj, A.At(i, j))
					sbj++
				}
				sbi++
			}
			det = det + (math.Pow(-1, float64(x)) * A.At(0, x) * determinant(subMatrix, n-1))
		}
	}
	return det
}

// SolveLU Нахождение решение методоа LU
func SolveLU(Lower, Upper, vec matrix.Matrix) matrix.Matrix {
	n := vec.Cols()

	// Ly = b
	y := matrix.ZeroMatrixFrom(vec)
	for i := 0; i < n; i++ {
		y.SetAt(0, i, vec.At(0, i))
		for j := 0; j < i; j++ {
			y.SetAt(0, i, y.At(0, i)-Lower.At(i, j)*y.At(0, j))
		}
	}

	// Ux = x
	x := matrix.ZeroMatrixFrom(vec)
	for i := n - 1; i >= 0; i-- {
		x.SetAt(0, i, y.At(0, i))
		for j := i + 1; j < n; j++ {
			temp := Upper.At(i, j) * x.At(0, j)
			x.SetAt(0, i, x.At(0, i)-temp)
		}
		x.SetAt(0, i, x.At(0, i)/Upper.At(i, i))
	}
	return x
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
			outragedRightSide.SetAt(0, i, outragedRightSide.At(0, i)+x.At(0, j)*A.At(i, j))
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
