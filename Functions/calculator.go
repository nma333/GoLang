package main

import "fmt"

//Multiplication of two numbers.
var num1 int
var num2 int

func calculator(num1 int, num2 int) {

	fmt.Printf("\nBasic airthematic operations of %d and %d", num1, num2)

	add := num1 + num2
	fmt.Printf("\nAddition of %d and %d is %d", num1, num2, add)

	diff := num1 - num2
	fmt.Printf("\ndifference of %d and %d is %d", num1, num2, diff)
	//fmt.Printf("Multiplication of %v and %v is %v\n", a, b, multi)

	multi := num1 * num2
	fmt.Printf("\nMultiplication of %d and %d is %d", num1, num2, multi)

	div := num1 / num2
	fmt.Printf("\nDivision of %d and %d is %d", num1, num2, div)
}

func main() {
	//prints(2, 3)

	calculator(10, 9)
}

OutPut:

        Basic airthematic operations of 10 and 9
        Addition of 10 and 9 is 19
        difference of 10 and 9 is 1
        Multiplication of 10 and 9 is 90
        Division of 10 and 9 is 1
