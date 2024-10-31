package mymath

func Sqrt(x float64) float64 {
	var z float64 = 0.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
