// kullanıcı giriş yapar ve adı,soyadı,yaşadığı şehir ekrana yazdırılırsın

package main

import "fmt"

func main() {

	var isim string
	var soyisim string
	var yasam string
	var sor int
	var sor1 int

	for {

		fmt.Println("başlatmak için 1'e,programı sonlandırmak için ise 2'e basınız: ")
		fmt.Scanln(&sor)

		if sor == 1 {
			fmt.Println(" -- kayıt alma --")
			fmt.Scanln(&isim)
			fmt.Scanln(&soyisim)
			fmt.Scanln(&yasam)

			fmt.Println("lütfen isim,soyisim ve yaşadığınız ülke giriniz")
			fmt.Println("kayıtı görmek için 1'e basınız")

			if sor1 == 1 {
				fmt.Println("şuanki kaydınız:", "isim:", isim, "soyisim:", soyisim, "yaşadığı yer:", yasam)

			} else {
				fmt.Println("hata")
			}

		} else if sor == 2 {
			fmt.Println("program sonlandırıldı...")
		} else {
			fmt.Print("hata")
		}

	}
}
