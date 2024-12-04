package main

import "fmt"

func main() {

	fmt.Println("Structs in go lang")


	yontan := User {"yehonatan","yonatan7777@gmail.com", true, 35 }
}


type User struct { 
	Name string
	Email string
	Status bool
	Age init
}