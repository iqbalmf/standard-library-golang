package main

import (
	"fmt"
	"time"
)

func main()  {
	var duration time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration - duration2

	fmt.Println(duration, duration2)
	fmt.Println(duration3)
	fmt.Printf("duration : %d", duration3)
}