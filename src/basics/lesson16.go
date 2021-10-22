/*
	Functions
	Golang'te Fonksyionlar first class nesnelerdir
	Tanımlanabilirler, fonksiyon içinde fonksiyon olabilir
	Anonim fonksiyonlar olabilir (closures)
	Fonksiyona fonksiyon parametre olarak geçilebilir
	Tip olarak tanımlanabilirler
	Slice ya da Map (value olarak) elemanı olabilirler
	Bir struct'un filed olabilirler
	Channel'larda send/receive parametresi olabilirler
*/

package main

import (
	"fmt"
	"log"
	"os"
)

type strFunc func(string) string

func main() {

	// by value olarak
	// numbers, bool, arrays, structs

	// by reference olarak
	// pointer, strings, slices, maps, channels

	// Parametre ya da argümanlar pass by value ya da passed by reference olarak alınabilir.
	s := "Hello World"
	passByValue(s)
	fmt.Printf("s -> %s\n", s)

	n := "Yemeksepeti"
	passByRef(&n)
	fmt.Printf("n -> %s\n", n)

	// Variadics -> N tane argüman
	greet("mert", "ismail")

	// Return Values
	// fonksiyonlar birden çok sonuç geri dönebilir

	result, err := returnValues("Hello Golang")
	if err == nil {
		fmt.Printf("With return values -> %s\n", result)
	}

	// Naked Return
	nakedResult := nakedReturn("Melin")
	fmt.Printf("naked return -> %s\n", nakedResult)

	// recursive function
	fact := fact(10)
	fmt.Printf("fact -> %d\n", fact)

	// Closure & Anomim & Scope Functions
	add := func(a, b int) int {
		return a + b
	}

	fmt.Printf("anonim func -> %d\n", add(5, 6))

	sc := scope()
	fmt.Printf("scope func -> %d\n", sc())

	func() {
		fmt.Println("Anonymous func is run")
	}() // anonymous func

	// Tip ya da argüman olarak Fonksiyonlar
	greet := func(name string) string {
		return "hello " + name
	}
	var fk strFunc
	fk = greet
	fmt.Printf("Type func -> %s\n", fk("Mert"))

	funcs := []func(){
		func() {
			fmt.Println("Func 1 is run")
		},
		func() {
			fmt.Println("Func 2 is run")
		},
		func() {
			fmt.Println("Func 3 is run")
		},
	}
	f := runManyFuncs(funcs)
	fmt.Println(f)
	f()

	// defer() fonks return etmeden önce çalışır
	defer func() {
		fmt.Println("exit - main")
	}()
	deferFunc()
	fmt.Println("after deferFunc")

	// asıl kullanım yeri
	createTempFile()

	a := 1
	defer fmt.Println("defer a", a)
	a = 100
	fmt.Println("a", a)

	fmt.Println("do -> ", do())
}

func passByValue(s string) {
	fmt.Println("Incoming", s)
	s = "Changed"
	fmt.Println("Changed", s)
}

func passByRef(s *string) {
	fmt.Println("Incoming", *s)
	*s = "Changed"
	fmt.Println("Changed", *s)
}

func greet(names ...string) {
	for _, v := range names {
		fmt.Println("Hello " + v)
	}
}

func returnValues(n string) (string, error) {
	return n, nil
}

func nakedReturn(n string) (result string) {
	result = "Hello " + n
	return
}

//recursive func
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// closure
func scope() func() int {
	outer_var := 2
	foo := func() int {
		return outer_var
	}
	return foo
}

func runManyFuncs(fList []func()) func() {
	return func() {
		for _, f := range fList {
			f()
		}
	}
}

func deferFunc() {
	defer func() {
		fmt.Println("exit - deferFunc")
	}()

	fmt.Println("Hello deferFunc")
}

func createTempFile() {
	f, err := os.Create("tmp\\foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func do() (a int) {
	defer func() {
		a = 300
	}()
	a = 1
	return
}
