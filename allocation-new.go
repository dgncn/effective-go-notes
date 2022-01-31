package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func main() {
	//go dilinde allocation new ve make ile yapılır.
	//new anahtar kelimesi memory allocation yapar ancak diğer dillerden
	//farklı olarak initialize yapmaz zero valued şekilde allocation yapar. belirtilen tipte
	//Mutex ve Buffer structları sıfır değerleri olarak veri yapısı döner
	//buffer boş buffer dönerken mutex unlock olmuş mutext döner

	y := new(T)
	y.name = "asd"
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println(y)
}

type T struct {
	F    os.File
	name string
	m    sync.Mutex
	b    bytes.Buffer
}

//output

/*
&{{<nil>} asd {0 0} {[] 0 0}}
*/
