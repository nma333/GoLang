//D-types in Golang.
package main

import "fmt"

func main() {
	fmt.Println(" Main different types of Data-Types on Golang:")
	//Integer
	var num int = 19
	fmt.Println("Integer:", num)

	//Float
	var double float32 = 89.90
	fmt.Println("Float:", double)

	//string.
	var name string = "Naveen"
	fmt.Println("Strings:", name)

	//boolean.
	var decision bool = true
	fmt.Println("Boolean:", decision)

	//undefined value.
	var class int
	fmt.Println("Undefined value is :", class)

}

Output:
        Main different types of Data-Types on Golang:
        Integer: 19
        Float: 89.9
        Strings: Naveen
        Boolean: true
        Undefined value is : 0
