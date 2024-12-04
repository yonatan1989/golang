package main

import "fmt"

func main (){

fmt.Println("Maps in golang")
lang := make(map[string]string)

	lang["JS"] = "JavaScript"
	lang["PY"] = "Python"
	lang["RB"] = "Ruby"

fmt.Println("List of all languages", lang)
fmt.Println("List of all PY :\n " ,lang["PY"] )

	delete (lang, "RB")

fmt.Println("List of all languages", lang)

for key, vaule := range lang {

	fmt.Printf("For key %v, value is %v\n", key, vaule)
}

}

