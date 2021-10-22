/*
	Flow Control & Execution Order

	Bir golang uygulaması kodu yukarıdan aşağıya doğru çalıştırarak çalışır. buna sequential executing ya da lineer executing denir
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	number := 100
	if number > 99 {
		fmt.Println("The number is greater then 99")
	}

	// short declaration
	if greet, err := greeting("Mert"); err == nil {
		fmt.Println(greet)
	}

	// short declarion with initializing
	if number = 200; number > 199 {
		fmt.Println("The number is greater then 199")
	}

	// switch - case
	os := "os"
	switch os {
	case "os":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux Geek")
	default:
		fmt.Println("Others")
	}

	// switch with identifier initialization
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS Hipseter")
	case "linux":
		fmt.Println("Linux Geek")
	default:
		fmt.Println("Others: ", os)
	}

	switch {
	case number > 199:
		fmt.Println("The number is greater then 199")
	case number == 199:
		fmt.Println("The number is equal 199")
	case number < 199:
		fmt.Println("The number is smaller then 199")
	}

	// label ve break
switchStatement:
	switch number {
	case 200:
		for i := 0; i < 5; i++ {
			break switchStatement
		}
		fmt.Println(200)
	default:
		fmt.Println("default case...")
	}

	// goto with label
Start:
	fmt.Println("number: ", number)
	if number > 210 {
		goto End
	} else {
		number++
		goto Start
	}
End:
}

func greeting(name string) (string, error) {
	return "Hello " + name, nil
}
