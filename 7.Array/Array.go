package main

import "fmt"

func main() {
	var num [5]float32
	num[0] = 1.2
	num[4] = 1.90
  fmt.Printl("The array elements are:")
	fmt.Println(num)

}

Output:
      The array elements are:
      [1.2 0 0 0 1.9]
