package main

import "fmt"

func main() {
	fmt.Println(len("new(int)"))

	/*var (
		x int
		y string
	)
	x,y =4, "ss"

	d:=7
	fmt.Println(d)
	fmt.Println(x,y)

	*/

}

/*
godoc : paket içerikleri hakkındaki dokumanları çıkartan go source kodlarını işletir

Every package should have a package comment
Her paketin paket yorumu olmalı

Names:
Bir paketin dışından bir adın görünebilirliği büyük harf ile başlaması veya başlamaması
seçeneğine göre değişir
	Getters:
		Dil bazında bir getter yok ihtiyaç da yok. kendin yazabilirsin demeye getirilmiş.

owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}

GetOwner yazma direkt Owner yaz kısa ve öz olsun demişler. Set tarafında SetOwner yazılabilir.
field ile method farkında metodu büyük fieldı küçük yazma öneriliyor.
zaten method baş harfi büyük olsun ki
exported names diye geçebilsin

InterfaceNames:
	Reader, Writer gibi er ile biterler. read, write gibi func. ifade edilir
	aynı anlamı ifade eden ve sistemde yer alan metotlar varsa, aynı adda kullanmayı öneriyor.
	String'e convert yapacaksan String yaz diyor ToString yapma.

To avoid confusion, don't give your method one of those names unless it has the same signature
and meaning. Conversely, if your type implements a method with the same meaning
as a method on a well-known type, give it the same name and signature; call your
string-converter method String not ToString.

Çoklu isim barındıracak kelimelerde _ yerine workMyCaps kullan demiş. PascalCase, camelCase gibi

Semicolons:
	Lexer kodu okurken en sona noktalı virgülü koyuyor sen ekleme.

lexer'ın noktalı virgül koydugu kelimeler: break continue fallthrough return ++ -- ) }

Hatta alttaki iki örnekteki farka bakılırsa 1. örnek kullanımına go izin vermez
if koşulundaki en son satıra süslü koyulmasının sebebi de lexer'ın kapatma parantezini
gördüğünde noktalı virgül koyması.
1-
if x < f(y)
{

2-
if x < f(y) {

// break loop

*/
