/*
	Struct
	Composite Type ailesinden bir tiptir
	Yapısal veri saklamanın en kısa halidir.
*/

package main

import "fmt"

type userInfo struct {
	firstName string
	lastName  string
	age       int
}

type userAnonymous struct {
	string
	int
}

type personInfo struct {
	name    string
	age     int
	address address
}

type personPromoted struct {
	name string
	age  int
	address
}

type address struct {
	city, country string
}

func main() {
	userInfo1 := userInfo{
		firstName: "Mert",
		lastName:  "Kimyonsen",
		age:       30,
	}

	// ya da direkt çıplak bir şekilde tanımlama yapabiliyoruz
	userInfo2 := userInfo{
		"Mert",
		"Kimyonsen",
		30,
	}

	fmt.Printf("userInfo1 -> %+v\n", userInfo1)
	fmt.Printf("userInfo2 -> %+v\n", userInfo2)

	// Anonymous Struct
	userInfo3 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Melin",
		lastName:  "Kimyonsen",
		age:       24,
	}
	fmt.Printf("userInfo3 -> %+v\n", userInfo3)

	userInfo4 := struct {
		firstName string
		lastName  string
		age       int
	}{
		"Hülya",
		"Kimyonsen",
		24,
	}
	fmt.Printf("userInfo3 -> %+v\n", userInfo4)

	// empty initialize
	userInfo5 := userInfo{}
	fmt.Printf("userInfo5 -> %v\n", userInfo5)
	fmt.Printf("userInfo5 -> %+v\n", userInfo5)

	// new ile object initialize
	userInfo6 := new(userInfo)
	fmt.Printf("userInfo6 -> %+v, type -> %[1]T\n", userInfo6) // hafızada userInfo tipi için rezerve et bana
	userInfo6.firstName = "Merve"
	fmt.Printf("userInfo6 -> %+v\n", *userInfo6)

	//new tüm tipler için geçerli ama &Type sadece structlar için geçerlidir.
	a := new(int)
	b := new(string)
	c := new(map[string]string)
	fmt.Printf("a, type: %T value: %[1]v *value: %v\n", a, *a)
	fmt.Printf("b, type: %T value: %[1]v *value: %v\n", b, *b)
	fmt.Printf("c, type: %T value: %[1]v *value: %v\n", c, *c)

	// Golang bize explicit derefernce imkanı sağlar
	// (*userInfo).firstName yerine -> userInfo.firstName
	fmt.Printf("%v\n", (*userInfo6).firstName)
	fmt.Printf("%v\n", userInfo6.firstName)

	var userAnonymous = userAnonymous{}
	fmt.Printf("userAnonymous -> %T\n", userAnonymous)
	userAnonymous.string = "Hello String"
	userAnonymous.int = 100
	fmt.Printf("userAnonymous -> %+v, type -> %[1]T\n", userAnonymous)

	// İç içe nested struct
	p1 := personInfo{
		name: "Mert",
		age:  30,
		address: address{
			city:    "Istanbul",
			country: "Turkey",
		},
	}

	fmt.Printf("personInfo -> %+v\n", p1)

	// Promoted Fields
	p2 := personPromoted{}
	p2.name = "Mert"
	p2.age = 30
	p2.address = address{
		city:    "Ankara",
		country: "Turkey",
	}
	fmt.Printf("person promoted -> %+v\n", p2)
	fmt.Printf("city promoted -> %+v\n", p2.city)
	fmt.Printf("country promoted -> %+v\n", p2.country)

	// Struct'lar değer tipli oldukları için karşılaştırılabilirler
	p3 := personInfo{"Mert", 30, address{city: "Istanbul", country: "Turkey"}}
	p4 := personInfo{"Mert", 30, address{city: "Istanbul", country: "Turkey"}}
	fmt.Printf("p3 equals p4 -> %t\n", p3 == p4)
}

func increment(n *int) {
	fmt.Printf("n -> %T\n", n)
	fmt.Printf("val -> %v, type -> %[1]T\n", *n)
	fmt.Printf("val -> %v, type -> %[1]T\n", &n)

	x := &n
	fmt.Printf("x | val -> %v, type -> %[1]T\n", x)

	**x = 1
	fmt.Printf("x 2 | val -> %v, type -> %[1]T\n", x)
	fmt.Printf("n | val -> %v, type -> %[1]T", *n)
}
