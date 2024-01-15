// Run file using 'go run file_name.go'
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	for x := 0; x < 5; x++ {
		fmt.Println("value of x is",x)
	}

	names := []string{"Bittu","Singh","Kingh"}
	for index, value := range names {
		value = "new value"
		fmt.Printf("The value at index %v is %v \n",index,value)
		value = "newest value"
		names[index] = "Raja"
	}

	for _, value := range names {
		fmt.Printf("The value is %v \n",value)
	}

	fmt.Println(names)
}