/*
Interfaces
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(&v)
}
