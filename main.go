package main

import (
	"fmt"
	"os"
	"time"
)

func foo() bool {
	fmt.Println("start a new line")
	fmt.Println("adsadsfasdfasfas")
	if t := true; !t {
		fmt.Println(t)
		return true
	}
	return false
}

func main() {
	var x = 10
	fmt.Printf("x_value: %d, x_addr: %p\n", x, &x)

	var px = &x
	fmt.Printf("px_value: %d, px_addr: %p\n", *px, px)

	time.Sleep(time.Second * 1)
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is the pointer p: %p\n", p)
	fmt.Printf("here is the string *p: %s\n", *p)
	fmt.Printf("here is the string s: %s\n", s)
	fmt.Print("aldfjasldjflasdjfl")
	for index := 0; index < 10; index++ {
		fmt.Println(index)
	}

	foo()

	// var nil_pointer *int = nil
	// *nil_pointer = 1
	fmt.Printf("%s", os.Getenv("PATH"))
	// fmt.Printf("%s", os.Getenv("GOPATH"))
}
