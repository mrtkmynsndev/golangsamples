/*
	Struct
*/

package main

import (
	"fmt"
	"math"
	"strings"
)

// strcut alanlardan oluşan bir nesnedir
type Vertex struct {
	X int
	Y int
}

var (
	v1      = Vertex{X: 1, Y: 2}
	v2      = Vertex{Y: 3}
	v3      = Vertex{}
	pVertex = &Vertex{4, 5}
)

func main() {
	v := Vertex{X: 10, Y: 20} //struct oluşturma
	fmt.Println(v)

	p := &v               // v değişkenin prointer adresini oluşturuyoruz. İçerisinde v değişkenin tutulduğu address var
	*&p.X = 20            // p bir pointer olduğu için pointer üzerinden değişken değeri set etmek için bu şekilde kullanıyoruz
	fmt.Printf("%p\n", p) // pointer i ekrana yazdırmak için %p formatı kullanılır
	fmt.Println(v)

	fmt.Println(v1, v2, v3, pVertex)

	// Arrays
	fmt.Println("Arrays...")
	var a [2]string // array tanımında length type solunda. Boyutu sonradan değiştirilemez
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	// array'den farkı dinamik boyutlu olması
	var s []int = primes[1:4]
	fmt.Println(s)

	// slice array'in referansıdır
	// slice değişen değerler ana aray içinde de değişir
	names := []string{"Mert", "Ayberk", "Can", "Özgür"}
	name1 := names[0:2]
	name2 := names[1:3]
	fmt.Println(name1, name2)
	name2[0] = "XXX"
	fmt.Println(name1, name2)
	fmt.Println(names)

	// Slice Literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}

	fmt.Println(s1)

	//Slice defaults
	s = s[1:4]
	fmt.Println(s)
	s = s[:2] // low bound 0 kabul edilir
	fmt.Println(s)
	s = s[1:] // high bound arrayin en son elemanı kabul edilir
	fmt.Println(s)

	printSlice(s)

	var sNil []int
	printSlice(sNil)
	if sNil == nil {
		fmt.Println("The slice is nil")
	}

	// Slice with make
	m := make([]int, 5) // make ile dinamik array oluşturuyoruz. Array capacity kadar ilk değerlerle oluşur
	printSlice(m)

	m = make([]int, 0, 5) // 3 parametre ile oluşan arrayın lengthi 0 olur
	printSlice(m)

	// Slice of Slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// apend ile array'e eleman ekleyebiliyoruz. array nil olsa bile
	var s2 []int
	printSlice(s2)
	s2 = append(s2, 0)
	printSlice(s2)

	s2 = append(s2, 1, 2, 3, 4) // birden fazla elemen ekleyebiliyoruz
	printSlice(s2)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// MAP
	fmt.Println("MAP........")
	var m1 map[string]Vertex // map tanımı

	m1 = make(map[string]Vertex) // make ile map oluşturabiliriz
	m1["First"] = Vertex{10, 20}
	fmt.Println(m1["First"])

	// map literals
	var m2 = map[string]Vertex{ // type vermeden struct tanımlayabiliriz
		"First":  {10, 20},
		"Second": {},
	}

	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["First"] = 10
	m3["Second"] = 20
	fmt.Println("The value:", m3["First"])
	fmt.Println("The value:", m3["Second"])
	delete(m3, "First") // map içinden eleman siliyoruz
	fmt.Println("The value:", m3["First"])

	v3, ok := m3["First"] // Eğer key map'in içindeyse ok true, değilse false döner
	// Eğer key map'in içindeyse v3 değerini, değilse 0 döner
	fmt.Println("The value:", v3, "Present?", ok)

	// Function Values
	fmt.Println("Function Values.......")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(str(func(s1 string) string {
		return strings.Split(s1, " ")[0]
	}))

	// function clousere

	pos, neg := adder(), adder()

	fmt.Println(pos(1), neg(-2))
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func str(fn func(s1 string) string) string {
	return fn("Hello World")
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func printSlice(s []int) {
	// slice length -> len(), capacity -> cap() ile hesaplanır
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
