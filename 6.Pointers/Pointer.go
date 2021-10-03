package main

import "fmt"

func main() {
	one := 123

	var ptr *int = &one
	fmt.Println("Value of a variable:", one)

	fmt.Println("Value of a ptr:", *ptr)

	fmt.Println("Addres of a ptr:", ptr)

	fmt.Println("New value is:", *ptr*2)
}

Output:

        Value of a variable: 123
        Value of a ptr: 123
        Addres of a ptr: 0xc0000a6058
        New value is: 246
