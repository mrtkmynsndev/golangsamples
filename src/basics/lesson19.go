/*
 	Interface

	 Golang'ın en önemli konularından biri
	 2 tür interface anlamı bulunuyor
	 	1. Tip olan interface (empty interface)
		2. Davranış tanımlayan interface
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type customType int

type Number int

func (n Number) Positive() bool {
	return n > 0
}

/* Kütüphanedeki Stringer Interface

type Stringer interface {
	String() string
}

*/
func (n Number) String() string { // Stringer -> Interface Satisfaction yaptık
	return fmt.Sprintf("%v", int(n))
}

func (n Number) Format(f fmt.State, verb rune) {
	val := n.String()
	if verb == 81 {
		val = "\"" + val + "\""
	}
	fmt.Fprint(f, val)
}

type DemoRW struct{}

func (d DemoRW) Read(p []byte) (n int, err error) {
	return 1, nil
}

func (d DemoRW) Write(p []byte) (n int, err error) {
	return 1, nil
}

var _ io.ReadWriter = (*DemoRW)(nil) // compile time proof, doesn't allocate

func main() {
	// empty interface as Value
	// eğer bir değişken interface{} olarak tanımlanmışsa bunun init değeri nil
	// Rob Pike (golang kurucusu) der ki; interface{} says nothing
	// empty interface hafızada point ettiği bir yer olmadığı için değer nill olur

	var i interface{}
	fmt.Printf("val -> %v, type -> %[1]T\n", i)

	greetInterface("Hello")
	greetInterface(200)
	u := struct{ name string }{name: "Mert"}
	greetInterface(u)

	// Type Checking
	// zorda kalmadıkça kullanma
	printByType(nil)
	printByType(1)
	printByType(5.5)
	printByType("Hello")
	printByType(strings.NewReader("Hello Reader"))
	var custType customType
	custType = 5
	printByType(custType)

	// Davranış olarak Kullanım (Behaviour)
	// interface aslında bir davranış sergiler
	// Name Convention olarak sonuna -er takısı alır
	// Golang tek abstract type'dir
	x := Number(5)
	fmt.Println(x)
	fmt.Printf("%d\n", x)
	fmt.Println("number is positive: ", x.Positive())

	y := Number(20)
	fmt.Printf("Number is: %Q\n", y)

	drw := DemoRW{}
	fmt.Println(checkInterfaceAsReadWriter(drw))
	fmt.Println(drw.Read([]byte("Hello")))

	// Accept interface as function argument
	s, err := readStream(strings.NewReader("Hello Stream"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("byte stream read -> ", s)

	// read from file
	f, err := os.Open("tmp/foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs, err := readStream(f)
	fmt.Println("file read -> ", fs)
}

func greetInterface(i interface{}) {
	fmt.Printf("val -> %v, type -> %[1]T\n", i)
}

func printByType(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println(j, " this is nil")
	case int:
		fmt.Println(j, " this is int")
	case string:
		fmt.Println(j, " this is string")
	case bool, rune:
		fmt.Println(j, " this is bool or rune")
	case io.Reader:
		fmt.Println(j, " this is io.Reader")
	case customType:
		fmt.Println(j, " this is customType")
	default:
		fmt.Printf("%v, no idea -> %[1]T\n", j)
	}

}

func checkInterfaceAsReadWriter(i interface{}) bool {
	_, ok := i.(io.ReadWriter)
	return ok
}

func readStream(r io.Reader) (string, error) {
	b := make([]byte, 1024)

	n, err := r.Read(b)
	if err != nil {
		return "", fmt.Errorf("read stream error %w", err)
	}

	return fmt.Sprintf("read %d bytes: %s (%v)", n, string(b), b[:]), nil
}
