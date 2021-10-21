/*
	Boolean
	true ve false değerleri için kullanılır, 1 bit integer olarak ifade edilir
*/

package main

import "fmt"

func main() {
	var result bool

	fmt.Printf("%t\n", result)

	result = true
	fmt.Printf("%v\n", result)
	fmt.Printf("%t\n", result)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || false)
	fmt.Println(false || true)

}
