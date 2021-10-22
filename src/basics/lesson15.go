/*
	Type Conversion
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)

	// Ascii to Int
	v := "10"
	if s, err := strconv.Atoi(v); err == nil { // Atoi ParseInt eşittir
		fmt.Printf("%T, %[1]v", s)
	}
}
