package main

import "fmt"

func main() {
	Hello := Student{"Naveen", 06, 10}

	//Output with fmt.Println
	fmt.Println("OutPut of Struct using Println:")
	fmt.Println(Hello)

	fmt.Printf("OutPut of Struct using Printf with +v:")
	fmt.Printf("%+v\n", Hello)

}

//Creating the Structure.
type Student struct {
	Name     string
	Roll_num int
	Marks    float32
}

Output:

          OutPut of Struct using Println:
          {Naveen 6 10}
          OutPut of Struct using Printf with +v:
          {Name:Naveen Roll_num:6 Marks:10}
