// Run file using 'go run file_name.go'
package main

import "fmt"

func main() {
	languages := make(map[string] string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["SQL"] = "Structured Query Language"

	fmt.Println("List of all languages is: ",languages)
	fmt.Println("JS languages is: ",languages["JS"])
	delete(languages,"SQL")
	fmt.Println("List of all languages is: ",languages)

	// loops in maps
	for _,value := range languages {
		fmt.Printf("For key v, the value is %v\n",value)
	}
}