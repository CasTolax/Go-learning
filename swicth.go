package main

import "fmt"

func main() {
	val := 5 // değer

	switch val { // switch - case if else bloğu gibidir,yapıyı kontrol eder
	case 1:
		fmt.Println("birinci")
	case 2:
		fmt.Println("ikinci")
	case 3:
		fmt.Println("üçüncü") // 10'a kadar devam eder
	default:
		fmt.Println("sayı yok")
	}
}
