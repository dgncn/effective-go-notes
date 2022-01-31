package main

import "fmt"

func main() {
	// for içerisinde kullanılan switch condition'ı içerisinde break + "label adı" yapıldığında
	// tanımlı label'a kadar kırar. for döngüsü Loop adında etiketlendi ve içeride etiketlenmiş
	// Loop'u kır dendi.
Loop:
	for i := 0; i < 5; i++ {
		fmt.Print(i)
		switch i {
		case 0, 2:
			fmt.Println("0test")
		case 3:
			break Loop
		}
	}
}
