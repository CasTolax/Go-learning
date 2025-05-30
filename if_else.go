// if-else-else if anlatım ve örnekleri

package main

import "fmt"

func main() { // sürücü belgesi örneği

	var yas int // yaş

	fmt.Println("lütfen yaşınıız giriniz: ")

	fmt.Scanln(&yas) // kullanıcı yaşını girer

	if yas >= 18 { // eğer yaşı 18'den büyük ise
		fmt.Println("sürücü belgesi için uygunsunuz")
	} else if yas <= 18 { // eğer yaşı 18'den küçük ise
		fmt.Println("maalesef uygun değilsiniz")
	} else { // istenen durum olmaz ise
		fmt.Println("hata:sadece isteneni yapınınız")
	}

}
