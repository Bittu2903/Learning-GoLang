// Run file using 'go run file_name.go'
package main

import (
	"fmt"
)

func main() {
	// Arrays
	var ages [3]int = [3]int{2,7,5}
	var names = [3]string{"bittu","singh","kingh"}
	courses := [3]string{"MCA","M.Tech"}
	fmt.Println(ages,names,courses)

	// Slices
	city := []string{"Delhi","HR","UP"}
	city = append(city,"MP")
	sliceOne := city[1:]
	fmt.Println(city,sliceOne)


}