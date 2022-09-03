

package main

import (
	"fmt"
	"time"
)

func main() {

	addTwoNumbers(888, 999)
	time.Sleep(1 * time.Second)
	datatypes()

}


func addTwoNumbers(x int, y int)  {
	/* basic function test */
	spacer()
	fmt.Println("addTwoNumbers(x int, y int):")
	total := 0
	total = x + y
	fmt.Println("sum for: ", x, " + ", y, " = ", total)
	spacer()
}

func spacer() {
	fmt.Println(" ")
}


func datatypes() {
	spacer()
	fmt.Println("datatypes():")
	// all numeric types default to 0

	// unsigned int with 8 bits
	// Can store: 0 to 255
	var myIntu8 uint8

	// signed int with 8 bits
	// Can store: -127 to 127
	var myInt8 int8

	// unsigned int with 16 bits
	// var myint uint16

	// signed int with 16 bits
	// var myint int16

	// unsigned int with 32 bits
	// var myint uint32

	// signed int with 32 bits
	// var myint int32

	// unsigned int with 64 bits
	// var myint uint64

	// signed int with 64 bits
	// var myint int64


	myIntu8 = 255
	fmt.Println(myIntu8)

	myInt8 = -127
	fmt.Println(myInt8)


	spacer()



}