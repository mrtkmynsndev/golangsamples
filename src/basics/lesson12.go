/*
	Pointer
*/

package main

import "fmt"

func main() {
	number := 10
	fmt.Printf("number-> %10v\n", &number)

	increment(number) // number'in value of fonksiyonu çalıştı
	fmt.Printf("number | value -> %v, value of -> %v\n", number, &number)

	incrementValueOf(&number)
	fmt.Printf("*number | value -> %v, value of -> %v\n", number, &number)

}

func increment(n int) { // golang değişkenin değerini geçti, buna by pass value deniyor
	n++
	fmt.Printf("n | value -> %v, value of -> %v\n", n, &n)
}

func incrementValueOf(n *int) { // number'in address of fonksiyonu çalıştı
	*n++
	fmt.Printf("*n | value -> %v, value of -> %v\n", *n, &n)
}
