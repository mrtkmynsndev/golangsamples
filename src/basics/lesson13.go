/*
	Struct Method and Pointer Reciever
*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// Pointer reciever
func (c Circle) area() float64 { // Circle nesnesinde herhangi bir değişiklik ihtiyaç olmadığı için reciever pointer kullanmadık
	return math.Pi * c.r * c.r
}

func (c *Circle) extend(r float64) { // Circle nesnesinde herhangi bir değişiklik yapmak istediğimiz zaman pointer kullanırız
	c.r = c.r * r
}

func main() {
	circle := Circle{0, 0, 4}
	fmt.Printf("Circle Area -> %v\n", area(circle))

	// Do this with pointer reciever
	circle2 := Circle{0, 0, 5}
	fmt.Printf("Circle Area Reciever -> %v\n", circle2.area())

	circle3 := Circle{0, 0, 5}
	fmt.Printf("circle3 are -> %v\n", circle3.area())
	circle3.extend(5)
	fmt.Printf("after extend circle3 are -> %v\n", circle3.area())

}

func area(c Circle) float64 {
	return math.Pi * c.r * c.r
}
