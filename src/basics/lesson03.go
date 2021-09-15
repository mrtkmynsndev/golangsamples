package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// pointers
	// pointer * operandı ile kullanılır
	// pointer, bir değerin bellekteki adresini tutar
	// & operandı ise pointer oluşturmak için kullanılır

	i := 50
	pointerOfI := &i // pointerOfI bir pointer oldu. içerisinde i değerinin tutulduğu address var
	fmt.Println("Value of i = ", i)
	fmt.Println("pointerOfI pointer address = ", pointerOfI)
	fmt.Println("value of pointer address = ", *pointerOfI)

	j := 32
	pointerOfJ := &j // pointerofJ pointer oldu. İçerisinde j değişkenin tutulduğu bellek addresi var
	k := pointerOfJ
	*k = 66
	fmt.Println("Value of j: ", j)
	fmt.Println("Pointer address of pointerOfJ: ", pointerOfJ)
	fmt.Println("Pointer address of k: ", k)
	fmt.Println("Value of pointerOfJ: ", *pointerOfJ)
	fmt.Println("Value of k: ", *k)

	// Pointer in functions
	fmt.Println("Pointer in function............")
	number := 50
	calculateWithoutPointer(number)
	fmt.Println("value of number in main: ", number)
	fmt.Println("With Pointer...")
	calculateWithPointer(&number)
	fmt.Println("value of number in main: ", number)

	fmt.Printf("main --> %p\n", initPerson())
}

func calculateWithoutPointer(number int) {
	number++
	fmt.Println("pointer address of number in function: ", &number)
	fmt.Println("value of number in function: ", number)
}

func calculateWithPointer(number *int) {
	*number++
	fmt.Println("pointer address of number in function: ", number)
	fmt.Println("value of number in function: ", *number)
}

func initPerson() *person {
	m := person{name: "Mert", age: 30}
	fmt.Printf("initPerson --> %p\n", &m)
	return &m
}
