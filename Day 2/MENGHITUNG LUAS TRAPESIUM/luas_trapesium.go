package main

import "fmt"

func main() {
	var a, b, h float64

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&h)

	luas := 0.5 * (a + b) * h

	fmt.Println("Luas trapesium adalah", luas)
}
