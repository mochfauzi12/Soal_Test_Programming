package utils

import "fmt"

func TukarNulai() {
	a, b := 10, 20
	fmt.Printf("Sebelum ditukar: a=%d, b=%d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("Setelah ditukar: a=%d, b=%d\n", a, b)
}
