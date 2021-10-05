package main

import "fmt"

func main() {
	languages := make(map[string]string)

	//Inserting the values.
	languages[".py"] = "Python"
	languages[".js"] = "JavaScript"
	languages[".java"] = "Java"
	languages[".go"] = "Go-lang"
	languages[".asm"] = "Assembly-Language"
	languages[".rb"] = "Ruby"

	fmt.Println("After deleting the Value pairs are:")
	for key, value := range languages {
		fmt.Printf("%v:%v\n", key, value)
	}

	delete(languages, ".asm")

}


Output:
        After deleting the Value pairs are:
        .py:Python
        .js:JavaScript
        .java:Java
        .go:Go-lang
        .asm:Assembly-Language
        .rb:Ruby
