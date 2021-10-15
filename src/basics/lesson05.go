/*
Methods with reciever type
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{5, 12}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	Scale2(&v, 10)
	fmt.Println(v.Abs())

	p := &Vertex{3, 4} // direkt vertex'in pointer oluşturduk
	ScaleFunc(p, 10)   // pointer adres değerini bekliyor
	fmt.Println(p.Abs())

	v2 := &Vertex{12, 13}
	fmt.Println(v2.Abs())
	AbsFunc(*v2) // fonksiyon pointer değerini bekliyor

}

// Vertex struct için tanımlanan özel yapı
// fonksiyon hemen yanına hangi struct için kullanacağımızı tanımlıyoruz
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
// struct'u pointer  ile sarmallayabiliyoruz. Bu sayede struct orginal değerini koruyoruz
// referans tipli örneği gibi
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer tipini fonksiyona parametre olarak gönderebiliyoruz
// etki alanı diğer pointerlar gibi aynı
func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
