/*
	nil
	nil -> Pointer, Interface, Map, Channel ve Fonksiyon tipleri için zero value (uninitialized) değeridir
	Tanımlanmamış ya da undefined bir state değildir

	nil'in tipi yoktur
	bir keyword değildir
	Golang ile gelen ön tanımlı yapıdır.
*/

package main

import "fmt"

type animal struct {
	types  []string
	skills map[string]string
}

func (a *animal) name() {
	fmt.Println("random name ....")
}

func (a *animal) printTypes() {
	if a == nil {
		fmt.Println("Nothing types ....")
	} else {
		for _, v := range a.types {
			fmt.Println("name: ", v)
		}
	}
}

func main() {
	// Hangi durum nil'dir?
	// Pointers: Hiçbir şeye point etmediğinde
	// Maps: Initialize edilmemişse
	// Slices: Destekleyen (underlying) bir array yoksa
	// Channels: Initialize edilmemişse
	// Functions: Initialize edilmemişse
	// Interfaces: Değer atanmamışlarsa hatta nil pointer olsalar bile

	var a1 *animal // not allocated or initialized! ** işaret eden bir pointer yok
	fmt.Printf("animal -> %v\n", a1)

	a1.name()

	a2 := animal{}
	fmt.Printf("animal -> %T %[1]v %[1]p\n", a2)
	fmt.Printf("animal types -> %T %[1]v %[1]p\n", a2.types)
	fmt.Println(a2.types == nil)
	fmt.Printf("animal skills -> %T %[1]v %[1]p\n", a2.skills)
	fmt.Println(a2.skills == nil)

	// nil Pointer vs Pointer
	var s *string
	var ss = new(string) // allocated!  artık zero-value sahip
	fmt.Printf("type: %T value: %[1]v\n", s)
	fmt.Printf("type: %T value: %[1]p pvalue: %v\n", ss, *ss)

	var a3 *animal
	a3.printTypes()

	a4 := animal{}
	a4.printTypes()
	a4.types = []string{"mert"}
	a4.printTypes()

}
