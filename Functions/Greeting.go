package main

import "fmt"

func main() {
	prints(1)
	greet()
}

func prints(num int) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d= %d\n", num, i, (num * i))

	}
}

func greet() {
	var name string = "Megha"
  fmt.Println("Good morning ",name)
}


OutPut:
          1 * 1= 1
          1 * 2= 2
          1 * 3= 3
          1 * 4= 4
          1 * 5= 5
          1 * 6= 6
          1 * 7= 7
          1 * 8= 8
          1 * 9= 9
          1 * 10= 10
          Megha
