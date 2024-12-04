package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main () {
	fmt.Println("switch in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(10) + 1

	fmt.Println("switch in golang", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 6:
		fmt.Println("You can move 6 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough		//  move to case 
	default:
		fmt.Println("You are default")	
	}

}