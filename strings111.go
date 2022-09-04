package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	
	fmt.Println(" ")
	fmt.Println("strings111.go")

	time.Sleep(1 * time.Second)

	f1()
	time.Sleep(1 * time.Second)
	f2();
}

func f1() {
	/*using Short Variable Declaration*/
	var string111 string
	string111 = "wonderful world"

	fmt.Println((string111))
	fmt.Println("type: ", reflect.TypeOf(string111))

	string222 := "always be happy"
	fmt.Println(string222)
	fmt.Println("type: ", reflect.TypeOf(string222))
}


func f2() {
	/* go uses lexical scoping*/
	x := true
	if x {
		y := 1
		if x != false {
			fmt.Println(x)
			fmt.Println(y)
		}
	}
}

/*

some notes
if the name of a variable begins with a lowercase letter, it can only be accessed within the current package.
this is considered at "unexported" variables


if the name of a variable begins with a capital letter, it can be accessed from
packages outside the current packag. This is considered "exported" variables

*/
