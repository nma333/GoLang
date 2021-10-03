package main

import (
	"fmt"
)

func main() {
	var num [5]float32
	num[0] = 1.2
	num[4] = 1.90
	fmt.Println(num)

	var names = [3]string{"Megha", "Naveen", "Manish"}
	fmt.Println("Array elements using for loop:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]=%v\n", i, names[i])
	}

}

Output:
        [1.2 0 0 0 1.9]
        Array elements using for loop:
        names[0]=Megha
        names[1]=Naveen
        names[2]=Manish
