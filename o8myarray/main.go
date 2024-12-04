package main

import "fmt"

func main() {

	fmt.Println("welcome to array")

	var fruitList [4]string
	
	fruitList[0] = "Apple"

	fruitList[1] = "Bannana"

	fruitList[2] = "Peach"

	fruitList[3] = "Apple"

	fmt.Println(fruitList)

	fmt.Println( len(fruitList))

	var vegList = [3]string{"Tommato", "beans", "mushroom"}


	fmt.Println(len(vegList))
}