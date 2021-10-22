/*
	Map
	key=value şeklinde veri tutabilen yapıdır.
	hash-table kullanır
	add, get ve delete O(1) constant-time complexity sahiptir
*/

package main

import (
	"fmt"
	"reflect"
)

var m map[string]string // initializer değeri nill'dir

func main() {
	fmt.Printf("value -> %v, len -> %d\n", m, len(m))

	// make ile instance oluşturabiliriz
	m2 := make(map[string]string, 5) // map literal oluşturuyoruz
	m2["firstName"] = "Mert"
	fmt.Printf("value -> %v, len -> %d\n", m, len(m))

	// nill map
	printMap(nil)

	// map için önden hafıza rezervasyonu yapılabilir
	m3 := make(map[string]int, 100)
	m4 := make(map[string]int)
	m3["yemeksepeti"] = 1_000_000
	m4["yemeksepeti"] = 1_000_000
	fmt.Printf("%p\n", m3)
	fmt.Printf("%p\n", m4)

	// reflection paketindeki DeepEqual ile mapler birbirleri ile karşılaştırılabilir
	fmt.Println(reflect.DeepEqual(m3, m4))

	m5 := map[string]string{
		"key1": "hello",
		"key2": "world",
	}
	fmt.Printf("m5 -> %v, len -> %d\n", m5, len(m5))

	v, ok := m5["key1"]
	if ok {
		fmt.Println("v", v)
	}

	m5["key3"] = "Mert"
	fmt.Printf("m5 -> %v, len -> %d\n", m5, len(m5))

	delete(m5, "key3")
	fmt.Printf("m5 -> %v, len -> %d\n", m5, len(m5))

	// map iteration
	for k, v := range m5 {
		fmt.Printf("key -> %s, value -> %s\n", k, v)
	}
}

func printMap(m map[string]string) {
	fmt.Printf("value -> %v, len -> %d\n", m, len(m))
}
