package main

import "fmt"

func main() {

	var password string //= "Naveen*333"

	fmt.Println("Enter your password:")
	fmt.Scanln(&password)

	if password == "Naveen*333" {
		fmt.Println("You entered the right password")

	}
}

Output:
        Enter your password:
        Naveen*333
        You entered the right password
