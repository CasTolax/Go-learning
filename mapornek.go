package main

import "fmt"

func main() {
	mevsimler := map[string]string{ // mevsimler listesi
		"yaz":      "sıcak",
		"kış":      "soğuk",
		"ilkbahar": "ılıman",
		"sonbahar": "serin",
	}

	val, ok := mevsimler["yaz"] // yaz adında anahtar varmı?
	fmt.Println(val, ok)        // çıktı "sıcak true"

	fmt.Println(len(mevsimler)) // kaç anahtar olduğunu yazdırır

	// fmt.Println("önce:", mevsimler) // normal hali
	// delete(mevsimler, "kış") // kış anahtarını sil
	// fmt.Println("sonra:", mevsimler) // yenilenmiş hali
}
