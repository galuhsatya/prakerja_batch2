package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	if isPrima(bilangan) {
		fmt.Printf("%d adalah bilangan prima\n", bilangan)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", bilangan)
	}
}

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
