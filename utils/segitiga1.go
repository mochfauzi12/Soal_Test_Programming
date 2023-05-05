package utils

import (
	"fmt"
)

// func main() {
// 	rows := 11
// 	fmt.Print("Jumlah Baris Adalah: 10")

func Segitiga1() {
	rows := 10
	//fmt.Print("Masukkan jumlah baris: ")
	//fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("Jumlah Baris Yang di Print adalah : 10 ")
}
