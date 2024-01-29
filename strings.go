package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Split("Iqbal Fauzan", " "))
	fmt.Println(strings.ToLower("Iqbal Fauzan"))
	fmt.Println(strings.ToUpper("Iqbal Fauzan"))
	fmt.Println(strings.Trim("    Iqbal Fauzan     ", " "))
	fmt.Println(strings.ReplaceAll("Iqbal Fauzan", "Iqbal", "Labqi"))
}
