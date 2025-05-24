package main

import "fmt"

func main() {
	var soru1 string
	fmt.Println("lütfen ismiinizi giriniz: ")

	n, err := fmt.Scanln(&soru1) // hata olcağını biliyorsak n,err yapılmalıdır

	if err != nil { // eğer sayı girilirse
		fmt.Println("sadece isminizi giriniz", n) // bunu yazdır
	} else { // eğer hata yoksa
		fmt.Println("merhaba", soru1) // bunu yazdır
	}

}
