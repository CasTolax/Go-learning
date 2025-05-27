package main

import "fmt"

func main() {
	var manav [5]string = [5]string{"elma", "armut", "çilek", "muz", "friut"}
	var kitap [5]string = [5]string{"1984", "yabancı", "sınırlar", "sokratesin savunması", "çiftlik"}
	var sor1 int
	var sor2 int

	for {
		fmt.Println(" başlatmak için 1'e çıkmak için ise 2 ye basınız")
		fmt.Scanln(&sor1)

		if sor1 == 1 {
			fmt.Println("başlatıldı")
			fmt.Println("1.liste için 1'e tıklayın 2.liste için 2'ye tıklayınız")
			fmt.Scanln(&sor2)

			if sor2 == 1 {
				fmt.Println(manav)
			} else if sor2 == 2 {
				fmt.Println(kitap)
			} else {
				fmt.Println("hata: sadece 1 veya 2 girebilirsiniz")
				break
			}

		} else if sor1 == 2 {
			fmt.Println("program sonlandırıldı")
			break
		} else {
			fmt.Println("hata: sadece 1 veya 2'ye basınız")
		}

	}
}
