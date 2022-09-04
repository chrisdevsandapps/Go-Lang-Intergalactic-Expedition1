
/*
functions are their own data types

*/
package main

import (
	"fmt"
)


func f1() {
	fmt.Println("f1():")
}

func f2(x int, y int) {
	sum := x + y
	fmt.Println("f2() sum: ", sum)
}


/* creating a function to another function */
func f3(inputFunc func(int, int) int) {

	fmt.Println("inputFunc: ", inputFunc(555, 333))
}




// closure
func returnFunc(x string) func() {

	emojis := "😀😀😀"
	
	return func() {
		fmt.Println("input value: ", x, emojis)
	}
}






func main() {

	// can assign a function to a variable:
	xxx1 := f1
	xxx1()



	// can assign a function to a variable even if has parameters:
	xxx2 := f2
	xxx2(888, 999)



	// pwede din anonymous function na agad yung sa variable
	xxx3 := func(x int, y int) {
		sum := x + y
		fmt.Println("sum: ", sum)
	}

	xxx3(888, 999)



	/* calling the function directly from where it is defined */
	xxx4 := func(x int) int {
		return x * -1
	}(8)

	fmt.Println("xxx4: ", xxx4)




	
	xxx5 := func(x int, y int) int {
		return x + y
	}
	f3(xxx5)

	xxx6 := func(x int, y int) int {
		return x - y
	}
	f3(xxx6)



	// immediately invoked function
	func() {
		fmt.Println("this is iife")
	}()




	// closure
	returnFunc("hahaha")()
	returnFunc("wazzup!")()

	// can also be stored in a variable:
	xxx7 := returnFunc("hehehe")
	xxx7()
}