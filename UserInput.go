package main

import "fmt"

func main() {
	var name string
	var id int
	fmt.Scanln(&name, &id)
	
	fmt.Println(name)
	fmt.Println(id)
}