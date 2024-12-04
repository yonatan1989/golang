package main

import (
	"fmt"
)


func main () {
	fmt.Println("loop in golang")

	string koko 
	

	days := []string{"Sunday", "wednesday", "Friday", "saturday", koko}

	fmt.Println(days)
	for i:=0; i < len(days); i++{
		fmt.Println(i)
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {

		fmt.Println("loop in golang", index, day)
	}
}