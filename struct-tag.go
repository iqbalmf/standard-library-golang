package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func main()  {
	sample := Sample{"Iqbalmfauzan13", "", ""}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")
	fmt.Println(required)
	fmt.Println(max)

	person := Sample{
		Name: "Iqbal",
		Address: "",
		Email: "12312",
	}
	fmt.Println(IsValid(person))
}

func IsValid(value any)(result bool){
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}