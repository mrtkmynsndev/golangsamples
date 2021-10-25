/*
	Errors
	error interface satify eden bir tiptir ve interface zero value'su nil dir.
	içinde Error() metodu bulunduran herhangi bir concrete type artık bir error olur

	Golang dilinde exception handling yoktur, error bir tiptir ve diğer tipler gibi yönetilmelidir
	error'lar görmezden gelinmemelidir
	Fail fast mantığıyla önce hatalar yakalanıp duruma göre ilerlenmelidir
	error'ü yanlızca bir kere handle edin

*/

package main

import (
	"errors"
	"fmt"
	"os"
)

type errType int

const (
	_ errType = iota // 0, for skip
	badRequest
	internal
	unkown
)

var errSpecial = errors.New("special error")

type MyError struct {
	kind  errType
	inner error
}

// error struct satify ediyoruz

func (e MyError) Error() string {
	switch e.kind {
	case badRequest:
		return "Bad Request"
	case internal:
		return "Oppss! something wrong!"
	}
	return "Unkown error"
}

func main() {

	var err error
	fmt.Printf("%v %[1]T\n", err)

	err = errors.New("new error")
	fmt.Printf("%v %[1]T\n", err)

	// CUSTOM ERROR TYPES
	var badRequest = MyError{kind: badRequest}
	fmt.Printf("badRequest -> %v, type -> %[1]T\n", badRequest)

	// short declaration
	if err := myErrorFactory(0); err != nil {
		fmt.Printf("short declaration of myError -> %v, type -> %[1]T\n", err)
	}

	err = myErrorFactory(0)
	fmt.Printf("myError -> %v, type -> %[1]T, message -> %[1]s\n", err)

	err = myErrorFactory(1)
	fmt.Printf("myError -> %v, type -> %[1]T, message -> %[1]s\n", err)

	err = myErrorFactory(2)
	fmt.Printf("myError -> %v, type -> %[1]T, message -> %[1]s\n", err)

	err = myErrorFactory(3)
	fmt.Printf("myError -> %v, type -> %[1]T, message -> %[1]s\n", err)

	// WRAPPING ERRORS
	if err = myErrorFactory(5); err != nil {
		fmt.Println(fmt.Errorf("wrapping error: %w\n", err))
	}

	// UNWRAP ERROR
	if err := fileChecker("not_here.txt"); err != nil {
		fmt.Println(err)
		if unwrapError := errors.Unwrap(err); unwrapError != nil {
			fmt.Println(unwrapError)
		}
	}

	// errors.Is
	// bir error zincirinde başka bir error olup olmadığını kontrol edebiliriz
	if err := checkErrorSpecial(true); err != nil {
		if errors.Is(err, errSpecial) {
			fmt.Println("Error Special")
		} else {
			fmt.Println("Default Error")
		}
	}

	// errors.As
	// errors.Is tersi gibi, error belli bir tipe mi ait?
	if err := checkErrorSpecial(true); err != nil {
		fmt.Printf("err -> %v, isErrorSpecial -> %t\n", err, errors.As(err, &errSpecial))
	}

	// panic ve recover
	// uygulamanın aksayan yönlerini, stack trace görmek için kullanılabilir ama gerçek uygulamalarda kullanmaktan kaçınmamız gereken bir fonksiyondur :)
	// Rop Pike der ki; Don't panic
	err = errors.New("panic error")
	if err != nil {
		panic(err)
	}

	// panic defer edilen fonksiyonların çalışmasını engellemez, nasıl mı?
	// tabi defer yaklaşımını önce taşımamız gerekiyor, aşağıdaki fonks panic'in önünde tanımlanmalıdır.
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover", p)
		}
	}()
}

func myErrorFactory(kind int) error {
	switch kind {
	case int(badRequest):
		return MyError{kind: badRequest}
	case int(internal):
		return MyError{kind: internal}
	default:
		return MyError{kind: unkown}
	}
}

func fileChecker(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("in fileChecker %w", err)
	}
	f.Close()
	return nil
}

func checkErrorSpecial(isSpecial bool) error {
	if isSpecial {
		return errSpecial
	} else {
		return errors.New("default error")
	}
}
