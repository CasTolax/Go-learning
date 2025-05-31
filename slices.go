package main

import "fmt"

func main() {

	var isimler [5]string = [5]string{"isim", "isim2", "isim3", "isim4"} // dizi
	dilim := isimler[2:]                                                 // 2'den itibaren devam et
	dilim[0] = "ahmet"                                                   // 0.index i ahmet yap

	fmt.Println(isimler) // yazdır
	fmt.Println(dilim)
	// slices sayesinde diziler daha esnek kullanılabilir
}
