package main

import "fmt"

func main() {
	languages := make(map[string]string)

	//Inserting the values.
	languages["py"] = "Python"
	languages["js"] = "JavaScript"
	languages["java"] = "Java"
	delete(languages, "py")

	for key, value := range languages {
		fmt.Printf("%v:%v\n", key, value)
	}

}

Output:
        java:Java
        py:Python
        js:JavaScript
