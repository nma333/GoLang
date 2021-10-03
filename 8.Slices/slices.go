package main

import "fmt"

func main() {
	//Step-1.
	var recipies = []string{"Eggrice", "Chicken-Biriyani"}
	fmt.Print("Recipies lis is as follows:")
	for i := 0; i < len(recipies); i++ {
		fmt.Printf("\nrecipies[%d] is %v", i, recipies[i])
	}

	//Adding items.
	recipies = append(recipies, "Mutton-Curry", "Bamboo-Biriyani")
	fmt.Print("Recipies list after adding items as follows:")
	for i := 0; i < len(recipies); i++ {
		fmt.Printf("\nrecipies[%d] is %v", i, recipies[i])
	}
}



OutPut:
        Recipies lis is as follows:
        recipies[0] is Eggrice
        recipies[1] is Chicken-Biriyani
