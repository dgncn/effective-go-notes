package main

import (
	"fmt"
)

func main() {
	//dönüş parametreleri isimlendirildiyse, fonk. başladığında
	//tipi ne ise o tip üstünden varsayılan(zero valued) değerlerini alırlar.
	//eğer fonk. değersiz şekilde direkt return derse, named parametreleri de varsayılan şekilde döner
	namedFuncExample("hey")
}

func namedFuncExample(s string) (n int, text string) {
	fmt.Printf("%v %v ", n, text)
	if text == "" {
		fmt.Println("text \"\" a eşittir ")
	}
	return
}
