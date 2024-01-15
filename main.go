// Run file using 'go run file_name.go'
// Creating basic server in GoLang
package main

import (
	"fmt"
	"io"
	"net/http"
)
func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from getHome")
	io.WriteString(w,"Hello from get Home function")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from getHello")
	io.WriteString(w,"Hello from get Hello function")
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from getIndex")
	io.WriteString(w,"Hello from get Index function")
}

func main() {
	http.HandleFunc("/",getHome)
	http.HandleFunc("/hello",getHello)
	http.HandleFunc("/index",getIndex) 
	err := http.ListenAndServe(":3333",nil)
	if err != nil{
		fmt.Println("The error occured is ",err)
	}
}