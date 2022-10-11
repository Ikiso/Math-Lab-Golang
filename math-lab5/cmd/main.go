package main

import (
	"errors"
	"fmt"
	"math-lab5/internal/methods"
	"math-lab5/pkg/matrix"
	"os"
)

func init() {
	// Согласно варианту 23
	accuracys := [6]float64{0.001, 0.0001, 0.00001, 0.000001, 0.0000001, 0.00000001}
	mat := [][]float64{
		{94, 0.68, 0.92, 3.23, 0.4},
		{0.68, 14, 0.38, 1.48, 8.8},
		{0.92, 0.38, 15, 1.28, 1.67},
		{3.23, 1.48, 1.28, 16, 0.95},
		{0.4, 8.8, 1.67, 0.95, 17},
	}
	size := len(mat)

	norm := make([]float64, size)
	lambda := 0.0
	iteration := 0
	println("\n\nСтепенной метод")
	for i := 0; i < len(accuracys); i++ {
		_, norm, lambda, iteration = methods.Stepwise(mat, size, accuracys[i])
		min, max := methods.FindMaxAndMin(mat, size, accuracys[i])
		printResult(norm, min, max, lambda, iteration, accuracys[i])
	}
	println("\n\nМетод произведений")
	for i := 0; i < len(accuracys); i++ {
		norm, lambda, iteration = methods.ScalarProducts(mat, size, accuracys[i])
		min, max := methods.FindMaxAndMin(mat, size, accuracys[i])
		printResult(norm, min, max, lambda, iteration, accuracys[i])
	}
	println("\n\nМетод Рэлея")
	for i := 0; i < len(accuracys); i++ {
		norm, lambda, iteration = methods.PrivateRaleigh(mat, size, accuracys[i])
		min, max := methods.FindMaxAndMin(mat, size, accuracys[i])
		printResult(norm, min, max, lambda, iteration, accuracys[i])
	}
}

func main() {

}

func printResult(result []float64, min, max, lambda float64, iteration int, eps float64) {
	fakeVector := matrix.ConvertVectorToMatrix(result)
	fmt.Printf("\nКол-во итераций: %d,Точность: %e, Искомое собственное число: %f, "+
		"Максимум: %f, Минимум: %f, Собственный вектор: %v",
		iteration, eps, lambda, max, min, fakeVector.String())
}

func write() (value int) {
	for {
		if value < 0 {
			err := errors.New("incorrect input value")
			println(err)

			_, err = fmt.Fscan(os.Stdin, &value)
			if err != nil {
				panic(err)
			}
		}

		break
	}
	return
}
