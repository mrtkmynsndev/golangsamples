/*
	Döngüler ve İterasyonlar

	Golang'de döngü ve iterasyonlar için for yapısı bulunur.
	do-while yapısı yoktur
*/

package main

import "fmt"

func main() {
	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("i -> %d\n", i)
	}

	// implicit control
	users := []string{"Hello", "World", "Golang"}

	for i, user := range users {
		fmt.Printf("index -> %d, val -> %s\n", i, user)
	}

	// index'i va ommit edebiliriz
	// blank identifier ediyoruz compile time'da görmezden gel anlamındadır
	for _, user := range users {
		fmt.Printf("user -> %s\n", user)
	}

	persons := map[string]string{
		"name": "Mert",
		"team": "Özgür",
		"mate": "Can",
	}

	for k, v := range persons {
		fmt.Printf("key -> %s, val -> %s\n", k, v)
	}

	// infinity loop
	i := 0
	for {
		i++
		fmt.Printf("i -> %d\n", i)
		if i > 5 {
			break
		}
	}

	// continue
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// statement control
	sum := 1
	for sum < 6 {
		fmt.Println("sum -> ", sum)
		sum += sum
	}

	// label kullanımı
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i -> %d, j -> %d\n", i, j)
			if j == 2 {
				break outer
			}
		}
	}
}
