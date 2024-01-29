package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main()  {

	//CONTAINER / LIST
	data := list.New()
	data.PushBack("Iqbal")
	data.PushBack("Muhammad")
	data.PushBack("Fauzan")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)
	next := head.Next()
	fmt.Println(next.Value)
	next = next.Next()
	fmt.Println(next.Value)
	//manual mendapatkan value

	for element := data.Front(); element != nil; element = element.Next()  {
		fmt.Println(element.Value)
	} //menggunakan looping mendapatkan value


	//CONTAINER / RING
	dataRing := ring.New(5)
	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = "Value "+ strconv.Itoa(i)
		dataRing = dataRing.Next()
	}
	dataRing.Do(func(value interface{}) {
		fmt.Println(value)
	})
}