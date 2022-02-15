package main

import "fmt"

func main() {
	//eğer x slice'ı y slice'ına atanırsa ikisi de aynı array referansını izler
	//bir fonksiyona parametr olarak geçen slice, çağıran tarafında slice'ın
	//elemanları gözükür durum oluşturur. pointer geçmek benzeri
	//slice tanımı
	slice := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(slice)
	//slice değişkenindeki 1. indexten başlayıp dizinin 2. elemanına kadar al
	//croppedSlice := slice[0:2]
	//fmt.Println(croppedSlice)
	//for ile her dönüşte her birini yeni bir slice'a yazmak

	for i := 0; i < 3; i++ {
		newSlice := slice[i:]
		newSlice2 := slice[i : i+1]
		fmt.Printf("slice[i:] : %v \n", newSlice)
		fmt.Printf("slice[i:i+1] : %v \n", newSlice2)
	}

	//slice birlestirme ornegi
	/*
		s1 := []int{1, 2, 5}
		s2 := []int{3}
		fmt.Println(SliceConcat(s1, s2))
	*/
}

//func (f *os.File) Read(buf []byte) []byte File package Read func.
/*
func SliceConcat(slice1 []int, slice2 []int) []int {
	//s3 := new([]int)
	l := len(slice1)
	s3 := slice1[0 : l+len(slice2)]
	return s3
}
*/
