package main

import (
	"fmt"
	"math"
	"os"
)

type Request struct {
	iter       int64
	accuracy   float64
	methodName string
	rootValue  float64
}

func main() {
	accuracys := [6]float64{0.001, 0.0001, 0.00001, 0.000001, 0.0000001, 0.00000001}
	conditional := true
	for {
		if !conditional {
			break
		}

		var a, b float64
		fmt.Println("Введите сначала первую точку диапзоноа затем, вторую" +
			"\nИсходные значения примера {-2.5; -2 , 1.5 ; 2}")

		_, err := fmt.Fscan(os.Stdin, &a)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Fscan(os.Stdin, &b)
		if err != nil {
			panic(err)
		}

		for _, accuracy := range accuracys {
			Call(a, b, accuracy)
		}

		fmt.Println("\nВы хотите снова ввести диапозон? [y]/[n]")

		var choice string

		_, err = fmt.Fscan(os.Stdin, &choice)
		if err != nil {
			panic(err)
		}

		switch choice {
		case "y":
			conditional = true
			continue
		case "n":
			conditional = false
			break
		}
	}

}

func Call(a, b, accuraccy float64) {
	r := new(Request)
	r.BisectionMethod(accuraccy, a, b)
	r.NewtonsMethod(accuraccy, a, b)
	r.ChordMethod(accuraccy, a, b)
	r.SimpleIterationMethod(accuraccy, a, b)

	fmt.Print("\n")
}

func f(x float64) float64 {
	return (math.Pow(3, x)) - 2*x - 5
}

func derivF(x float64) float64 {
	return math.Log(3)*math.Pow(3, x) - 2
}

func secondDerivF(x float64) float64 {
	return math.Pow(math.Log(3), 2) * math.Pow(3, x)
}

func getResultEqualEquation(x float64) float64 {
	return (math.Pow(3, x) - 5) / 2
}

func getResultDeriv(x float64) float64 {
	return (math.Log(3) * math.Pow(3, x)) / 2
}

func (r *Request) BisectionMethod(accuracy, a, b float64) {
	r.accuracy = accuracy
	r.iter = 1
	r.methodName = "Метод бисекции"
	r.rootValue = a

	for {
		ra := math.Abs(b - a)
		conditional := ra > r.accuracy
		if !conditional {
			break
		}
		r.rootValue = (a + b) / 2

		if f(r.rootValue) == 0 {
			break
		}

		if f(a)*f(r.rootValue) < 0 {
			b = r.rootValue
			r.iter++
		} else {
			a = r.rootValue
			r.iter++
		}
	}

	fmt.Printf("\nКол-во итераций: %d,Точность: %e,Имя метода: %s,Значение корня: %.9f",
		r.iter, r.accuracy, r.methodName, r.rootValue)
}

func (r *Request) NewtonsMethod(accuracy float64, a, b float64) {
	r.accuracy = accuracy
	r.methodName = "Метод Ньютона"
	r.rootValue = a
	r.iter = 1

	if f(a)*secondDerivF(a) <= 0 {
		r.rootValue = b
	}

	for {
		xi := r.rootValue - (f(r.rootValue) / derivF(r.rootValue))
		isLess := math.Abs(xi-r.rootValue) <= r.accuracy
		r.rootValue = xi
		if isLess {
			break
		} else {
			r.iter++
		}
	}

	fmt.Printf("\nКол-во итераций: %d,Точность: %e,Имя метода: %s,Значение корня: %.9f",
		r.iter, r.accuracy, r.methodName, r.rootValue)
}

func (r *Request) ChordMethod(accuracy float64, a, b float64) {
	r.accuracy = accuracy
	r.methodName = "Метод Хорд"
	r.iter = 1
	xi := b

	for {
		r.rootValue = b - ((b-a)/(f(b)-f(a)))*f(b)

		isLess := math.Abs(r.rootValue-xi) <= accuracy
		if isLess {
			break
		}

		cond := f(a)*f(r.rootValue) < 0
		if cond {
			b = r.rootValue
			xi = r.rootValue
		} else {
			a = r.rootValue
			xi = r.rootValue
		}
		r.iter++
	}

	fmt.Printf("\nКол-во итераций: %d,Точность: %e,Имя метода: %s,Значение корня: %.9f",
		r.iter, r.accuracy, r.methodName, r.rootValue)
}

func (r *Request) SimpleIterationMethod(accuracy float64, a, b float64) {

	r.accuracy = accuracy
	r.methodName = "Метод простых итераций"
	r.iter = 1
	r.rootValue = a
	max := math.Max(a, b)

	if math.Abs(getResultDeriv(math.Max(a, b))) < 1 {
		r.rootValue = (a + b) / 2
		for {
			xi1 := getResultEqualEquation(r.rootValue)

			isLessE := math.Abs(xi1-r.rootValue) > r.accuracy
			if !isLessE {
				r.rootValue = xi1
				break
			} else {
				r.rootValue = xi1
				r.iter++
			}
		}
	} else {
		r.rootValue = (a + b) / 2
		for {

			xi := r.rootValue - (f(r.rootValue) / derivF(max))
			isLess := math.Abs(xi-r.rootValue) <= r.accuracy
			r.rootValue = xi
			if isLess {
				break
			} else {
				r.iter++
			}
		}
	}

	fmt.Printf("\nКол-во итераций: %d,Точность: %e,Имя метода: %s,Значение корня: %.9f",
		r.iter, r.accuracy, r.methodName, r.rootValue)

}
