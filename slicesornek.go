package main

import "fmt"

// liste olucak
//slices ile kullanıcının belirlediği indexler seçilip değiştirebilecek

func main() {
	var kullanıcı1 int
	var baslat int
	var ekle string

	isimler := []string{"sayı", "merhaba", "selam"}
	for {
		fmt.Println("başlatmak için 1'e çıkmak için ise 2'ye basınız")
		fmt.Scanln(&baslat)

		if baslat == 1 {
			fmt.Println(" -- dizi kontrol --")
			fmt.Println("listeyi göster:1 listeye ekle:2 programı sonlandır:3")
			fmt.Scanln(&kullanıcı1)

			if kullanıcı1 == 1 {
				fmt.Println(isimler)
			} else if kullanıcı1 == 2 {
				fmt.Println("lütfen eklenicek metini giriniz: ")
				fmt.Scanln(&ekle)
				isimler = append(isimler, ekle)

				fmt.Println("liste:", isimler)
			} else if kullanıcı1 == 3 {
				fmt.Println("program sonlandırıldı...")
				break
			} else {
				fmt.Println("hata")
				break
			}

		} else if baslat == 2 {
			fmt.Println("program sonlandırıldı...")
			break
		}

	}

}
