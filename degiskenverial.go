package main

import "fmt"

func main() { // kullanıcıdan veri alınarak ekrana veri tipini yazdırmaya çalıştırmak
	var kullanıcı int
	var kullanıcı1 int
	var kullanıcı2 int

	for { // döngü
		fmt.Println("Uygulamayı başlatmak için 1'e, kapatmak için ise 2'ye basınız:")
		fmt.Scanln(&kullanıcı)

		if kullanıcı == 1 { // eğer kullanıcı 1'e basarsa çalıştır
			fmt.Println("Program başlatıldı.")
			fmt.Println("Sayı giriniz: ")
			fmt.Scanln(&kullanıcı1) // veri al

			fmt.Println("Veri alındı. Girdiğiniz sayının tipini görmek için 1'e basınız:")
			fmt.Scanln(&kullanıcı2) // veri al

			if kullanıcı2 == 1 { // eğer 1'e basılırsa çalıştır
				fmt.Printf("Girdiğiniz verinin tipi: %T\n", kullanıcı1) // ekrana veri tipini göster
			} else { // hata var ise ekrana bunu yazdır
				fmt.Println("Hata: Sadece 1'e basınız.")
			}
		} else if kullanıcı == 2 { // eğer 2'e basılırsa programı kapat
			fmt.Println("Program kapatıldı...")
			break // döngüyü kır
		} else { // eğer yanlış girilir ise bunu yazdır
			fmt.Println("Hata: Sadece 1 veya 2'ye basabilirsiniz.")
		}
	}
}
