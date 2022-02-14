package main

import "fmt"

func main() {
	//make'in new metodundan farkı zero olmayan yeni bir instance döner
	//sadece array,channel ve slice için örnek yaratır
	//sadece bu tiplerde işlenmesinin sebebi, bu veri yapılarının kullanılmadan
	//önce mutlaka initialize edilmesi zorunluluğudur
	mySlice := make([]int, 4, 50) // 50 intlik bir dizi allocate edilir, ardından 4 uzunluklu
	// slice create edilir
	mySecondSlice := new([]int) // verilen tipte yeni pointer döner zeroed slice.
	fmt.Println(mySlice, mySecondSlice, *mySecondSlice == nil)
	//output: [0 0 0 0] &[] true
}
