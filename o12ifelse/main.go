package main

import "fmt"


func main () {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	}else {
		result = "Exactly login count"
	}
	fmt.Println(result)
	
	if 9%2 ==0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

	if num :=3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is NOT less than 10")
	}

	// if err != nill {

	// }
	

}