package main

import (
	"fmt"
	"strconv"
)
func main()  {
	boolean, err:= strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}

	resInt, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resInt)
	}
}