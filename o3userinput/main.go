package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	//reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user for Start:")

	//comma or || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type is %T \n", input)

}
