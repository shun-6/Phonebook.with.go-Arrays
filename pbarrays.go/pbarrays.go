package main

import (
	"fmt"
)

type telkayıt struct {
	isim   string
	Numara string
}

func main() {
	var telrehber [5]telkayıt
	kayıtsayı := 0

	for {
		var secim int
		fmt.Println("1. Kişi Ekle")
		fmt.Println("2. Kişileri Listele")
		fmt.Print("Seçim Yapınız: ")
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			if kayıtsayı < len(telrehber) {
				var yenikayit telkayıt
				fmt.Print("Ad Giriniz: ")
				fmt.Scanln(&yenikayit.isim)
				fmt.Print("Numara Giriniz: ")
				fmt.Scanln(&yenikayit.Numara)
				telrehber[kayıtsayı] = yenikayit
				kayıtsayı++
			} else {
				fmt.Print("Yeterli Alana Sahip değilsiniz.")
			}

		case 2:
			fmt.Println("Telefon rehberi:")
			for i := 0; i < kayıtsayı; i++ {
				fmt.Println("İsim:", telrehber[i].isim, "\tNumara", telrehber[i].Numara)
			}
		}
	}
}
