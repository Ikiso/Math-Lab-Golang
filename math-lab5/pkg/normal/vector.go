package normal

import (
	"math"
)

func GetEuclideanNorm(vector []float64) float64 {
	sum := 0.0
	for i := 0; i < len(vector); i++ {
		sum += vector[i]
	}
	return math.Pow(sum, 0.5)
}
