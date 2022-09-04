

package main

import (
	"fmt"
)

func f1() {
	/* typical function */
	fmt.Println("f1():")
}

func addTwoNum(x int, y int) int {
	/* function with parameters, and return*/
	fmt.Println("addTwoNum():")
	return x + y
}



func addAndSub1(x int, y int) (int, int) {
	/* function with multiple return */
	return x + y, x - y
}



func addAndSub2(x int, y int) (z1 int, z2 int) {
	/* 
	function with multiple return, and named return 
	also, with 'defet keyword' 
	*/

	defer fmt.Println("hello i am from defer")
	z1 = x + y
	z2 = x - y
	fmt.Println("hello i am from before return")
	return;
}




func main() {

	f1()
	
	sum := addTwoNum(888, 999)
	fmt.Println("sum: ", sum)

	add1, subt1 := addAndSub1(888, 999)
	fmt.Println("addResult: ", add1, ", subractResult: ", subt1)

	add2, subt2 := addAndSub2(333, 222)
	fmt.Println("addResult: ", add2, ", subractResult: ", subt2)
	
}