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

	var string111 string
	string111 = "wonderful world"


	fmt.Println((string111))
	fmt.Println("type: ", reflect.TypeOf(string111))

	string222 := "always be happy"
	fmt.Println(string222)
	fmt.Println("type: ", reflect.TypeOf(string222))

}