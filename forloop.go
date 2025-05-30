// for anlatım ve örnekleri

package main

import "fmt"

func main() { // sayı girme örneği

	var sayı int // değişken

	for { // for döngüsü (Go da while döngüsünün yerine geçer)
		fmt.Println("sayı giriniz: ")

		fmt.Scanln(&sayı)

		// for ile ekrana yazdırma
		for i := 0; i < 5; i++ {
			fmt.Println("merhaba", i)
		}

	}
}
