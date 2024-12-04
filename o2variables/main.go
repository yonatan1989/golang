package main

import "fmt"

const LoginToken string = "ghhhakdsdfdsf" //Public

var jwtToken int = 30000

func main() {
	var username string = `kokofromTlv`
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallvar uint8 = 255
	fmt.Println(smallvar)
	fmt.Printf("Variable is of type: %T \n", smallvar)

	var smallFloat float32 = 255.455554444
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type

	var website = "learngo.in"
	fmt.Println(website)

	//no var style

	numberofusers := 30000
	fmt.Println(numberofusers)
	fmt.Printf("Variable is of type: %T \n", numberofusers)
	fmt.Println(LoginToken)

}
