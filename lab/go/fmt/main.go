package main

import "fmt"

var userName string

func main() {
	fmt.Println("give me your username")
	fmt.Scan(&userName)
	fmt.Printf("Username is %v", userName)
}
