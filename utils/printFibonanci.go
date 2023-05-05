package utils

import "fmt"

func PrintFibonacci(z int) {
	first, second := 0, 1
	fmt.Printf("%d %d ", first, second)
	for i := 3; i <= z; i++ {
		next := first + second
		fmt.Printf("%d ", next)
		first = second
		second = next
	}
}
