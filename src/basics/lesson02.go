/*  FlowControl */

package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //for döngüsü parantezler yok
		sum += i
	}
	fmt.Println("First for, The sum is: ", sum)

	for sum < 100 {
		sum += sum
	}
	fmt.Println("Second for, The sum is: ", sum)

	// infinity loop
	for {
		sum += 1
		if sum == 190 {
			fmt.Println("Infinity for: ", sum)
			break
		}
	}

	if sum > 189 { // if statement parantez yok
		fmt.Println("The sum is bigger than 189")
	}

	if sum := sum + 10; sum > 199 { // short if statement,
		fmt.Println("The sum is bigger than 199")
	}

	if sum := math.Pow(2, 4); sum > 16 {
		fmt.Println("The sum is bigger than 16")
	} else {
		fmt.Println("The sum is smaller than 16")
	}

	// switch statements
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	//defer
	defer fmt.Println("hello") // erteleme, en yakınındaki fonksiyon çağrıldıktan sonra execute olur
	fmt.Println("world")

	// defer ile execute edilen fonksyionlar stack atılıyor.
	// stack'ten last-in-first-out ile çağrılıyor
	fmt.Println("Counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Counting done.")
}
