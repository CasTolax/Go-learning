package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"elma":  5,
		"armut": 6,
		"vişne": 7,
	}
	fmt.Println(mp) // map böyle yapılabilir

	dosya := make(map[string]int)
	dosya["elma"] = 3
	dosya["armut"] = 7
	dosya["vişne"] = 4

	fmt.Println(dosya) // bir farklı hali
	//delete()
	//make()

}
