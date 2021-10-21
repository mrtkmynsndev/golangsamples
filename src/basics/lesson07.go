/*
	String
	İçinde unicode karakterleri bulundurun karakterler dizisidir.
	"" (çift tırnak) ya da ``(back-tick) içinde kullanılır
	Immutable yapıdadır yani değeri değişmez

*/

package main

import "fmt"

var message1 = "Hello World"
var message2 string

func main() {
	fmt.Println("message1: ", message1)
	fmt.Println("message2: ", message2)

	// Unicode ile birlikte
	weather := "Hava 42\u00B0 derece!"
	fmt.Println("Weather: ", weather)

	// unit8 ile byte
	// int32 ile rune aynı şeydir
	lastName := "Kimyonşen"
	fmt.Printf("%T %[1]v\n", lastName)
	fmt.Printf("%T %[1]v\n", []rune(lastName))
	fmt.Printf("%T %[1]v", []byte(lastName)) // ş 2 karakter oldupu için

	// Immmutable dedik o yüzden değiştirilemezler
	fmt.Println(lastName[1])
	// lastName[1] = "T" <- Hata: cannot assign to lastName[1]

	// String içindeki her bir karaktere erişmek için
	for index := range weather {
		// %c -> karakter
		// %d -> digit/sayı
		// %x -> hexadecimal/16'lık sistem
		fmt.Printf("%c | %d | $%x\n", weather[index], weather[index], weather[index])
	}

	// Slicing ile string değerleri biçimlendirebiliriz.
	fmt.Println(weather)
	fmt.Println(weather[:5])  // 0'dan 5'e kadar ama 5 hariç
	fmt.Println(weather[10:]) // 10'ten string sonuna kadar
	fmt.Println(weather[5:8]) // 5'den 8'ya kadar ama 6 hariç

	// String Concatenation ile strinleri toplamak mümkün
	firstName := "Mert"
	middleName := "Hasan"
	fmt.Println("String Concatenation -> ", firstName+" "+middleName)

	// String Literals -> 2 tür metinsel ifade yöntemi var
	// 1. Raw String: işlenmemiş anlamında, back-tick içinde kullanımı `` | not: içinde back-slash işlemez \n
	// 2. Interpreted String yani "" çift tırkan içinde, yorunlanmış anlamında
	rawString := `Hello \n Go Programming Language`
	interpretedString := "Hello \n Go Programming Language"
	fmt.Println("Raw String -> ", rawString)
	fmt.Println("Interpreted String -> ", interpretedString)

	rawString2 := `\n
					\n`
	fmt.Println("Raw String 2 -> ", rawString2)

}
