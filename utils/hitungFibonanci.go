package utils

func HitungFibonanci(z int) int {
	if z <= 1 {
		return z
	}

	first, second := 0, 1
	for i := 2; i <= z; i++ {
		temp := first + second
		first = second
		second = temp

	}
	return second
}
