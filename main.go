// Run file using 'go run file_name.go'
package main

import "fmt"

func main() {
	age := "23"
	name := "Bittu"
	floatedAge := 23.89
	// Println
	fmt.Println("Hello World!")  // automatically add next line \n
	
	// Print
	fmt.Print("first line \n")
	fmt.Print("Second line \n")
	fmt.Print("The age is ",age," and the name is ",name,"\n")

	// Printf (Formatter String)
	fmt.Printf("The format string with age %v and name %v \n",age, name)
	fmt.Printf("The format string with age %q and name %q \n",age, name)
	fmt.Printf("The format string with age %f and name %v \n",floatedAge, name)

	// Sprintf (save formatted string)
	str := fmt.Sprintf("The format string with age %f and name %v \n",floatedAge, name)
	fmt.Print(str)
}