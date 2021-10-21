/*
	Numerics
	Integers -> Tam Sayılar
	Float -> Ondalık Sayılar
*/

package main

import "fmt"

func main() {

	// complex türleri
	c1 := complex(5, 7)
	c2 := 2 + 5i
	fmt.Printf("c1 -> %v\n", c1)
	fmt.Printf("c2 -> %v\n", c2)

	// type conversion
	fmt.Println("------- Type Conversion -------")
	a := 42
	fmt.Printf("a -> %v %[1]T\n", a) // reuse tekniği

	b := 1.5
	fmt.Printf("b --> %v %[1]T\n", b)
	sum1 := a + int(b) // 2 sayıyı toplamak için b türünü int türüne çevirdik
	fmt.Printf("a + int(b) --> %v %[1]T\n", sum1)

	sum2 := float64(a) + b
	fmt.Printf("float64(a) + b --> %v %[1]T\n", sum2)

}
