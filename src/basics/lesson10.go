/*
	Slice and Array
	Composite Types yani birleşik tipler olarak geçer
	Array ve Slice birbirleri ile kardeş iki kavramdır, aralarında ufak ama önlemli bir far vardır!
*/

package main

import "fmt"

type user struct {
	firstName string
}

func main() {

	// Array -> içinde ayni türlerin olduğu ve boyutunun belli olduğu koleksiyonlardır.
	// kullanımı -> [boyut]Tip
	var a [3]int
	fmt.Printf("Val -> %v, Type -> %[1]T\n", a)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Printf("Val -> %v, Type -> %[1]T\n", a)

	// dizinin buyutu almak için -> len()
	fmt.Printf("a legth -> %v\n", len(a))

	// kısa aray tanımı
	b := [3]int{1, 2, 3}
	fmt.Printf("b -> %v, %[1]T\n", b)

	// boyutunu derleyici bulsun istiyorsak
	c := [...]string{"mert", "murat", "burak"}
	fmt.Printf("c -> %v, %[1]T\n", c)

	// araylar kopyalanır
	var furits = [...]string{"apple", "banana", "melon"}

	otherFurits := furits
	otherFurits[0] = "orange"
	fmt.Printf("otherFurits -> %v, len -> %d, cap -> %d\n", otherFurits, len(otherFurits), cap(otherFurits))
	fmt.Printf("furits -> %v, len -> %d, cap -> %d\n", furits, len(furits), cap(furits))

	// Slice dinamik ve ölçeklenebilir bir arraydir. Array in sabit boyutu olurken Slice olmayabilir buda bize esneklik sağlar
	users := []string{"Mert", "Melin"}
	users = append(users, "Merve", "Hülya")
	fmt.Printf("users -> %v, type -> %[1]T\n", users)

	// struct tanımlayarak array oluşturma
	var userList []user
	fmt.Printf("%v, type -> %[1]T\n", userList)
	userList = append(userList, user{"Mert"})
	fmt.Printf("%v, type -> %[1]T\n", userList)

	// Slice pointer kullanır kopyalanmaz ve içerde 3 şey saklar
	// 1. pointer
	// 2. length
	// 3. capacity

	s := make([]byte, 5)
	fmt.Printf("s -> %v, len -> %d, cap -> %d\n", s, len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("s -> %v, len -> %d, cap -> %d\n", s, len(s), cap(s))
	s = append(s, 2, 3, 4, 5)
	fmt.Printf("s -> %v, len -> %d, cap -> %d\n", s, len(s), cap(s))

	s2 := make([]string, 0, 4)
	fmt.Printf("s2 -> %v, len -> %d, cap -> %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, "Mert", "Melin", "Murat", "Merve")
	fmt.Printf("s2 -> %v, len -> %d, cap -> %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, "Hülya")
	fmt.Printf("s2 -> %v, len -> %d, cap -> %d\n", s2, len(s2), cap(s2))

	// slice end ve start vererek başka bir slice yaratabiliriz
	userSlice := users[1:3]
	userSlice[0] = "XXX"
	fmt.Printf("userSlice -> %v, len -> %d, cap -> %d\n", userSlice, len(userSlice), cap(userSlice))
	fmt.Printf("users -> %v, len -> %d, cap -> %d\n", users, len(users), cap(users)) // slice pointer olduğu için buradaki değer de değişti

	// three index slicing
	userThree := []string{"Foo", "Bar", "Baz"}
	userListWithCapacity := userThree[0:1:1] // start from index:0, end:1, cap:1
	fmt.Printf("userListWithCapacity -> %v, len -> %d, cap -> %d\n", userListWithCapacity, len(userListWithCapacity), cap(userListWithCapacity))

}
