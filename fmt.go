package main

import "fmt"

func main()  {
	firstName:= "Iqbal"
	lastName := "Fauzan"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}