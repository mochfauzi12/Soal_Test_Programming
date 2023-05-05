package utils

import "fmt"

func GanjilGenap() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d adalah bilangan genap\n", i)
		} else {
			fmt.Printf("%d adalah bilangan ganjil\n", i)
		}
	}
}
