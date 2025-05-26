package main

import "fmt"

func main() {
	var kullanıcı int
	var kullanıcı1 int
	var kullanıcı2 int

	for {
		fmt.Println("Uygulamayı başlatmak için 1'e, kapatmak için ise 2'ye basınız:")
		fmt.Scanln(&kullanıcı)

		if kullanıcı == 1 {
			fmt.Println("Program başlatıldı.")
			fmt.Println("Sayı giriniz: ")
			fmt.Scanln(&kullanıcı1)

			fmt.Println("Veri alındı. Girdiğiniz sayının tipini görmek için 1'e basınız:")
			fmt.Scanln(&kullanıcı2)

			if kullanıcı2 == 1 {
				fmt.Printf("Girdiğiniz verinin tipi: %T\n", kullanıcı1)
			} else {
				fmt.Println("Hata: Sadece 1'e basınız.")
			}
		} else if kullanıcı == 2 {
			fmt.Println("Program kapatıldı...")
			break
		} else {
			fmt.Println("Hata: Sadece 1 veya 2'ye basabilirsiniz.")
		}
	}
}
