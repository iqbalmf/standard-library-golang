package main

import (
	"fmt"
	"reflect"
)

type Sampe struct {
	Name, Address, Email string
	Age                  int
}
type Person struct {
	Name string
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func main() {
	readField(Sampe{"Iqbal", "", "", 0})
	readField(Person{"Iqbal"})
}
