package main

import "fmt"

func main() {

	fmt.Println("Welcome to o7myponiters")

	// var ptr *int

	// fmt.Println("Value of ptr ", ptr)

	mynum := 23

	var ptr = &mynum

	fmt.Println("vaule of mynumis ", ptr)
	fmt.Println("vaule of mynumis ", *ptr)

	*ptr = *ptr * 2 

	fmt.Println("vaule of mynumis ", mynum)


}
