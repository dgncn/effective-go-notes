package main

import "fmt"

func main() {
	/*
		arraylarin c'den farkı
		bir arrayin başka bir arraya atanması demek tüm değerlerinin kopyalanması demek
		bir fonksiyona array geçildiğinde fonk. dizinin kopyasını alır diziyi işaret eden pointer'ı değil
		arrayin size'ı tipin özelliğidir. yani 2 arrayin uzunluğu farklı ise ikisinin tipi farklıdır
		the types [10]int and [20]int are distinct.

	*/

	myCustomArray := [5]float64{123.123, 1232.12, 12312.23423, 234234.11, 443.213221}
	fmt.Println(myCustomArray)
	fmt.Printf("total: %v ", Sum(&myCustomArray))
}

func Sum(flArray *[5]float64) (sum float64) {
	for _, val := range *flArray {
		sum += val
	}
	return
}
