package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		my_num := 34
		your_num := 67
		fmt.Println(my_num + your_num)

	*/
	fmt.Println("The generated random number is:")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))

	//Time := time.Now()

	//fmt.Println(Time)

}

Output:
        The generated random number generates is:
        1
