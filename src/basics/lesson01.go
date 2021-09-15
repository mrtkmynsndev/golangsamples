package main

//projede kullancağımız paketleri tanımlıyoruz
import (
	"fmt"
	"math"
	"math/rand"
)

// constant can be character, string, boolean, or numeric values.
// constant := syntax ile değer ataması yapılamaz!
const (
	pi   = 3.14
	ToBe = true
)

//birden fazla değer tanımı
var (
	izmir    = 35
	istanbul = 34
)

// var ile değişkenler listesini tanımlayabiliriz
// aynı türü paylaşan değişkenlerdin tipini sonda tanımlayabiliriz
var c, python, java bool

// değişkenlere ilk değer ataması
var i, j int = 1, 2

func main() { //entry point
	fmt.Println("Hello from golang :)", rand.Intn(10))
	fmt.Println(add(5, 10))
	fmt.Println(multy(5, 4))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(10))

	var i int // önce değişken adı sonra tipi
	fmt.Println(i, c, python, java)

	fmt.Println(i, j)

	k := 3 // yeni bir değer ataması, dinamik değer ataması
	fmt.Println(k)

	fmt.Printf("İzmir Value: %v\n", izmir)
	fmt.Printf("İstanbul Value: %v\n", istanbul)

	var r float32 = 2.9
	var area = math.Pi * ((r / 2) * (r / 2))

	fmt.Printf("Area of Circle: %v\n", area)

	//type conversion
	num := 50
	f := float64(num)
	u := uint(f)
	fmt.Println(num, f, u)

	// Constant
	fmt.Println("Constant: ", pi)
	fmt.Println("Constant: ", ToBe)
}

// type variable isimden sonra geliyor
func add(x int, y int) int {
	return x + y
}

// parametreler aynı tipte ise en sondakine verilebilir
func multy(x, y int) int {
	return x * y
}

// fonksiyon birden fazla parametre dönebilir
func swap(x, y string) (string, string) {
	return y, x
}

// fonksiyon return tiplerine isim verilip değer ataması verilebilir
// argumansız return ler nacked return
// nacked return kısa fonksiyonda kullanılmalı
func split(sum int) (x, y int) {
	x = sum * 10
	y = sum - 10
	return
}
