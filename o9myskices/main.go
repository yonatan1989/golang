package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("welcome to o8myarray")

	var fruitList = []string{"Apple", "tomato", "Peach", "koko"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango")

	fruitList = append(fruitList[1:5])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores [0] = 234
   
	highScores [1] = 945
	
	highScores [2] = 456

	highScores [3] = 777

	highScores = append(highScores, 555, 66, 321)

	sort.Ints(highScores)

	fmt.Println(highScores)

	// how to remove a vaule from slices based on index 

	var  couress = []string{"java","javascript", "swift", "python" }
	var index int = 1
	couress = append(couress[:index], couress[index+1:]...)
	fmt.Println(couress)
}
