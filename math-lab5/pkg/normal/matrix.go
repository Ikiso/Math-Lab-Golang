package normal

import (
	"github.com/oelmekki/matrix"
	"math"
)

func FirstNormal(A matrix.Matrix, rows, cols int) float64 {
	sum := 0.0
	subSum := 0.0
	for i := 0; i < rows; i++ {
		subSum = 0.0
		for j := 0; j < cols; j++ {
			subSum += math.Abs(A.At(i, j))
		}
		if subSum > sum {
			sum = subSum
		}
	}
	return sum
}

func SecondNormal(A matrix.Matrix, rows, cols int) float64 {
	sum := 0.0
	subSum := 0.0
	for j := 0; j < rows; j++ {
		subSum = 0.0
		for i := 0; i < cols; i++ {
			subSum += math.Abs(A.At(i, j))
		}
		if subSum > sum {
			sum = subSum
		}
	}
	return sum
}

func ThirdNormal(A matrix.Matrix, rows, cols int) float64 {
	sum := 0.0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			sum += A.At(i, j) * A.At(i, j)
		}
	}
	return math.Sqrt(sum)
}
